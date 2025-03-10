package storage

import (
	"os"
	"testing"
	"vns/internal/models"

	"github.com/stretchr/testify/assert"
)

// TestStorePayloads tests the storing of payloads into payloads.json
func TestStorePayloads(t *testing.T) {
	// Prepare mock data for testing
	payloads := []models.Payload{
		{
			ID:             "CVE-2024-1234",
			Severity:       "HIGH",
			CVSS:           8.5,
			Status:         "fixed",
			PackageName:    "openssl",
			CurrentVersion: "1.1.1t-r0",
			FixedVersion:   "1.1.1u-r0",
			Description:    "Buffer overflow vulnerability in OpenSSL",
			PublishedDate:  "2024-01-15T00:00:00Z",
			Link:           "https://nvd.nist.gov/vuln/detail/CVE-2024-1234",
			RiskFactors:    []string{"Remote Code Execution", "High CVSS Score", "Public Exploit Available"},
		},
	}

	// Store the payloads in the file
	err := StorePayloads(payloads)
	assert.Nil(t, err)

	// Verify that the file exists and has the expected contents
	file, err := os.ReadFile("./payloads.json")
	assert.Nil(t, err)
	assert.Contains(t, string(file), "CVE-2024-1234")
}

// TestQueryPayloads tests the querying of payloads from payloads.json
func TestQueryPayloads(t *testing.T) {
	// Prepare mock data for testing
	payloads := []models.Payload{
		{
			ID:             "CVE-2024-1234",
			Severity:       "HIGH",
			CVSS:           8.5,
			Status:         "fixed",
			PackageName:    "openssl",
			CurrentVersion: "1.1.1t-r0",
			FixedVersion:   "1.1.1u-r0",
			Description:    "Buffer overflow vulnerability in OpenSSL",
			PublishedDate:  "2024-01-15T00:00:00Z",
			Link:           "https://nvd.nist.gov/vuln/detail/CVE-2024-1234",
			RiskFactors:    []string{"Remote Code Execution", "High CVSS Score", "Public Exploit Available"},
		},
	}
	StorePayloads(payloads)

	// Mock the query filters
	filters := map[string]string{
		"severity": "HIGH",
	}

	// Query the payloads
	result, err := QueryPayloads(filters)
	assert.Nil(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "CVE-2024-1234", result[0].ID)
}
