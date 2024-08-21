package handlers

import (
	"backend/kafka"
	"backend/logger"
	"backend/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func TrackEvent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var event models.Event
	err = json.Unmarshal(body, &event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eventData, err := json.Marshal(event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Logger.Printf("Event: %v", event.Type)
	kafka.ProduceMessage(event.Type, event.UserID, eventData)
	w.WriteHeader(http.StatusNoContent)
}
