package query

import (
	"encoding/json"
	"net/http"
	"vns/internal/models"
)

type QueryRequest struct {
	Filters map[string]string `json:"filters"`
}

// Modify HandleQuery to accept a function for querying payloads. This allows us to mock it during testing.
func HandleQuery(w http.ResponseWriter, r *http.Request, queryFunc func(map[string]string) ([]models.Payload, error)) {
	var queryRequest QueryRequest
	err := json.NewDecoder(r.Body).Decode(&queryRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query the database with filters using the provided queryFunc
	payloads, err := queryFunc(queryRequest.Filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payloads)
}
