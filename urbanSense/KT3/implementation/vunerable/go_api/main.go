package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SensorPayload map[string]any

func ingestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// RANJIVO: nema limita veličine body-ja (unbounded request body)
	// RANJIVO: parsing potencijalno troši CPU i memoriju za velike payload-e

	var payload SensorPayload
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	log.Printf("[INGEST] received keys=%d\n", len(payload))
	fmt.Fprintln(w, "ok")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ingest", ingestHandler)

	// RANJIVO: ListenAndServe bez timeout-a i bez dodatnih limita
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}