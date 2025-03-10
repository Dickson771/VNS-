package main

import (
	"log"
	"net/http"
	"vns/internal/github" // Import github to pass FetchFile function
	"vns/internal/query"
	"vns/internal/scan"
	"vns/internal/storage" // Import storage to use QueryPayloads

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the server...")

	// Initialize the router
	r := mux.NewRouter()

	// Define the /scan endpoint for processing the scanning requests
	// Wrap HandleScan in an anonymous function that injects fetchFile
	r.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) {
		// Call HandleScan and pass the fetchFile function
		scan.HandleScan(w, r, github.FetchFile)
	}).Methods("POST")

	// Define the /query endpoint for querying stored data
	// Pass storage.QueryPayloads to HandleQuery as the queryFunc argument
	r.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		query.HandleQuery(w, r, storage.QueryPayloads) // Pass storage.QueryPayloads as queryFunc
	}).Methods("POST")

	// Start the HTTP server
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r)) // Start the server on port 8080
}
