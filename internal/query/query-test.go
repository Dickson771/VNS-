package query

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"vns/internal/models"

	"github.com/stretchr/testify/assert"
)

// Mock query function to simulate querying payloads
func mockQueryPayloads(filters map[string]string) ([]models.Payload, error) {
	// Mock data to return when queried
	return []models.Payload{
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
	}, nil
}

// Test the HandleQuery function with a mock query function
func TestHandleQuery(t *testing.T) {
	// Prepare the request body
	requestData := map[string]interface{}{
		"filters": map[string]string{
			"severity": "HIGH", // Query filter
		},
	}
	requestBody, _ := json.Marshal(requestData)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Call HandleQuery with the mock query function
	HandleQuery(rr, req, mockQueryPayloads)

	// Assert that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert the response body contains the expected data
	var payloads []models.Payload
	err := json.Unmarshal(rr.Body.Bytes(), &payloads)
	assert.Nil(t, err)
	assert.Len(t, payloads, 1)
	assert.Equal(t, "CVE-2024-1234", payloads[0].ID)
	assert.Equal(t, "HIGH", payloads[0].Severity)
	assert.Equal(t, 8.5, payloads[0].CVSS)
	assert.Equal(t, "fixed", payloads[0].Status)
}
