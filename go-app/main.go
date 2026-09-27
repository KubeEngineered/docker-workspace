package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type AppStatus struct {
	Version     string    `json:"version"`
	Environment string    `json:"environment"`
	Status      string    `json:"status"`
	Uptime      string    `json:"uptime"`
	Timestamp   time.Time `json:"timestamp"`
}

var startTime time.Time

func main() {
	startTime = time.Now()

	// Retrieve port from ENV or fall back to default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "v3.0.0"
	}

	// Main root endpoint returning structured JSON
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		status := AppStatus{
			Version:     appVersion,
			Environment: getEnvOrDefault("APP_ENV", "local-dev"),
			Status:      "operational",
			Uptime:      time.Since(startTime).Round(time.Second).String(),
			Timestamp:   time.Now().UTC(),
		}

		json.NewEncoder(w).Encode(status)
	})

	// Dedicated health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"UP","healthy":true}`))
	})

	log.Printf("Server listening on port :%s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
