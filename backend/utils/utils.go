package utils

import (
	"sync"
	"time"
)

var userSessions = make(map[string]time.Time)
var mu sync.Mutex

func CalculateDuration(userID string, timestamp string) int {
	mu.Lock()
	defer mu.Unlock()

	// Parse the timestamp
	eventTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return 0
	}

	startTime, exists := userSessions[userID]
	if !exists {
		// If no session exists, start a new one
		userSessions[userID] = eventTime
		return 0
	}

	duration := int(eventTime.Sub(startTime).Seconds())
	// Reset the session start time
	userSessions[userID] = eventTime
	return duration
}
