package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type AppResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1.0.1"
	}

	// Root Endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request from %s on path %s", r.RemoteAddr, r.URL.Path)
		
		w.Header().Set("Content-Type", "application/json")
		resp := AppResponse{
			Status:  "success",
			Message: "Hello from updated Go App running inside Docker!",
			Version: version,
		}
		json.NewEncoder(w).Encode(resp)
	})

	// Health Check Endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"UP"}`)
	})

	log.Printf("Server starting on port %s (Version: %s)...", port, version)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
