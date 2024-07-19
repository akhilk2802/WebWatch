package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Event struct {
	EventType string                 `json:"eventType"`
	Data      map[string]interface{} `json:"data"`
	TimeStamp string                 `json:"timestamp"`
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process the event (e.g., send it to Kafka or store it in a database)
	log.Printf("Received event: %v", event)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event received"))
}

func main() {

}
