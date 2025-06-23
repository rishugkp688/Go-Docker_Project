package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct for JSON response
type Response struct {
	Message string `json:"message"`
}

// Handler for GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Welcome to the Go Docker App!"}
	respondJSON(w, http.StatusOK, resp)
}

// Handler for GET /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "OK"}
	respondJSON(w, http.StatusOK, resp)
}

// Handler for GET /greet?name=YourName
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	message := fmt.Sprintf("Hello, %s!", name)
	resp := Response{Message: message}
	respondJSON(w, http.StatusOK, resp)
}

// JSON response helper
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Main entry point
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/greet", greetHandler)

	port := ":8080"
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
