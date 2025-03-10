package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"vns/internal/models"
)

// StorePayloads stores the payloads into the JSON file (prevents duplicates)
func StorePayloads(payloads []models.Payload) error {
	// Define the file path where we store the data
	filePath := "./payloads.json"

	// Read the existing data from the file
	var existingPayloads []models.Payload
	data, err := os.ReadFile(filePath) // Replaced ioutil.ReadFile with os.ReadFile
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// If the file is not empty, unmarshal the data into existingPayloads
	if len(data) > 0 {
		err = json.Unmarshal(data, &existingPayloads)
		if err != nil {
			return fmt.Errorf("failed to unmarshal existing data: %v", err)
		}
	}

	// Create a map to track unique payloads by their ID
	payloadMap := make(map[string]models.Payload)

	// Add existing payloads to the map to ensure uniqueness
	for _, payload := range existingPayloads {
		payloadMap[payload.ID] = payload
	}

	// Add new payloads to the map (duplicates will be overwritten)
	for _, newPayload := range payloads {
		payloadMap[newPayload.ID] = newPayload
	}

	// Convert the map back into a slice
	var updatedPayloads []models.Payload
	for _, payload := range payloadMap {
		updatedPayloads = append(updatedPayloads, payload)
	}

	// Marshal the combined data back into JSON
	updatedData, err := json.Marshal(updatedPayloads)
	if err != nil {
		return fmt.Errorf("failed to marshal updated data: %v", err)
	}

	// Write the updated data back to the JSON file
	err = os.WriteFile(filePath, updatedData, 0644) // Replaced ioutil.WriteFile with os.WriteFile
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	log.Println("Successfully stored payloads in payloads.json")
	return nil
}

// QueryPayloads queries the payloads from the JSON file based on the filters
func QueryPayloads(filters map[string]string) ([]models.Payload, error) {
	// Define the file path for the payloads file
	filePath := "./payloads.json"

	// Read the data from the file
	data, err := os.ReadFile(filePath) // Replaced ioutil.ReadFile with os.ReadFile
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Unmarshal the data into a slice of Payloads
	var allPayloads []models.Payload
	err = json.Unmarshal(data, &allPayloads)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %v", err)
	}

	// Filter the payloads based on the severity filter
	var filteredPayloads []models.Payload
	for _, p := range allPayloads {
		if severity, ok := filters["severity"]; ok && p.Severity == severity {
			filteredPayloads = append(filteredPayloads, p)
		}
	}

	log.Printf("Found %d payload(s) with severity %s", len(filteredPayloads), filters["severity"])
	return filteredPayloads, nil
}
