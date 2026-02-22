package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type SensorPayload map[string]any

const maxBodyBytes = 1 * 1024 * 1024 // 1MB limit za demonstraciju

func ingestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// MITIGACIJA 1: limit veličine request body-ja
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	var payload SensorPayload
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&payload); err != nil {
		// Ako je prekoračen limit, Go vraća MaxBytesError
		var mbe *http.MaxBytesError
		if errors.As(err, &mbe) {
			http.Error(w, "request too large", http.StatusRequestEntityTooLarge) // 413
			return
		}
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	log.Printf("[INGEST] received keys=%d\n", len(payload))
	fmt.Fprintln(w, "ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ingest", ingestHandler)

	// MITIGACIJA 2: eksplicitni timeout-i na HTTP serveru
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	log.Println("Listening on :8080")
	log.Fatal(srv.ListenAndServe())
}