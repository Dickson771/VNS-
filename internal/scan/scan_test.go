package scan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"vns/internal/models"

	"github.com/stretchr/testify/assert"
)

// Mock fetchFile function to simulate fetching content from GitHub
func mockFetchFile(repo string, file string) ([]byte, error) {
	// Simulating a mock response from GitHub
	mockData := `[
		{
			"scanResults": {
				"vulnerabilities": [
					{
						"id": "CVE-2024-1234",
						"severity": "HIGH",
						"cvss": 8.5,
						"status": "fixed",
						"package_name": "openssl",
						"current_version": "1.1.1t-r0",
						"fixed_version": "1.1.1u-r0",
						"description": "Buffer overflow vulnerability in OpenSSL",
						"published_date": "2024-01-15T00:00:00Z",
						"link": "https://nvd.nist.gov/vuln/detail/CVE-2024-1234",
						"risk_factors": [
							"Remote Code Execution",
							"High CVSS Score",
							"Public Exploit Available"
						]
					}
				]
			}
		}
	]`
	return []byte(mockData), nil
}

// Test the HandleScan function with a valid request
func TestHandleScan(t *testing.T) {
	// Mock data for the request
	requestData := map[string]interface{}{
		"repo":  "velancio/vulnerability_scans",
		"files": []string{"vulnscan16.json"},
	}
	// Convert request data to JSON
	requestBody, _ := json.Marshal(requestData)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/scan", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Recorder to capture the response
	rr := httptest.NewRecorder()

	// Call HandleScan with the mock fetchFile function
	HandleScan(rr, req, mockFetchFile)

	// Assert that the status code is OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Debugging: Print the response body to understand the issue
	fmt.Println("Response Body:", rr.Body.String())

	// Decode the response body into a slice of Payloads
	var storedPayloads []models.Payload
	err := json.Unmarshal(rr.Body.Bytes(), &storedPayloads)
	assert.Nil(t, err)

	// Check that there is exactly 1 payload
	assert.Len(t, storedPayloads, 1)

	// Check the properties of the payload
	assert.Equal(t, "CVE-2024-1234", storedPayloads[0].ID)
	assert.Equal(t, "HIGH", storedPayloads[0].Severity)
	assert.Equal(t, 8.5, storedPayloads[0].CVSS)
	assert.Equal(t, "fixed", storedPayloads[0].Status)
}
