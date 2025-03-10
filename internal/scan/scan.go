package scan

import (
	"encoding/json"
	"log"
	"net/http"
	"vns/internal/models"
	"vns/internal/storage"
)

// HandleScan handles the POST request for the /scan endpoint
func HandleScan(w http.ResponseWriter, r *http.Request, fetchFile func(repo string, file string) ([]byte, error)) {
	var scanRequest struct {
		Repo  string   `json:"repo"`
		Files []string `json:"files"`
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&scanRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payloads []models.Payload
	for _, file := range scanRequest.Files {
		// Fetch the file content from GitHub (using the injected fetchFile function)
		content, err := fetchFile(scanRequest.Repo, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the JSON array from the file
		var scanData []struct {
			ScanResults struct {
				Vulnerabilities []models.Payload `json:"vulnerabilities"`
			} `json:"scanResults"`
		}

		// Unmarshal the file content into scanData (array of scan results)
		err = json.Unmarshal(content, &scanData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Extract the vulnerabilities and append them to the main list
		for _, scan := range scanData {
			payloads = append(payloads, scan.ScanResults.Vulnerabilities...)
		}
	}

	// Ensure the payloads slice is not empty before storing
	if len(payloads) == 0 {
		http.Error(w, "No vulnerabilities found", http.StatusInternalServerError)
		return
	}

	// Store the payloads in the JSON file
	err = storage.StorePayloads(payloads)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the stored payloads in the response body
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(payloads)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Successfully stored the payloads.")
	w.WriteHeader(http.StatusOK)
}

// fetchAndProcessFile fetches a file from the GitHub repository and processes it into Payloads.
/*func fetchAndProcessFile(repo, file string) ([]models.Payload, error) {
	// Fetch the content of the file from the GitHub repository
	content, err := github.FetchFile(repo, file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON content into Payload objects
	var payloads []models.Payload
	err = json.Unmarshal(content, &payloads)
	if err != nil {
		return nil, err
	}

	// Add scan metadata (like the current time) to each Payload
	for i := range payloads {
		payloads[i].PublishedDate = time.Now().Format(time.RFC3339)
	}

	return payloads, nil
}
*/
