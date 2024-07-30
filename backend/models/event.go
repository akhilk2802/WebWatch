package models

type Event struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	Referrer  string `json:"referrer,omitempty"`
	X         int    `json:"x,omitempty"`
	Y         int    `json:"y,omitempty"`
	Target    string `json:"target,omitempty"`
	UserID    string `json:"userId"`
	Timestamp string `json:"timestamp"`
	Duration  int    `json:"duration,omitempty"` // in seconds
}
