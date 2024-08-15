package models

import "time"

type Event struct {
	Type            string    `json:"type"`                        // Type of event: pageview, click, duration, scroll, mousemove, hover, form_submission, etc.
	URL             string    `json:"url,omitempty"`               // The URL where the event occurred
	Referrer        string    `json:"referrer,omitempty"`          // Referrer URL (for pageview events)
	UserID          string    `json:"userId"`                      // Unique identifier for the user
	Timestamp       time.Time `json:"timestamp"`                   // Timestamp when the event occurred
	X               int       `json:"x,omitempty"`                 // X-coordinate (for click, mousemove events)
	Y               int       `json:"y,omitempty"`                 // Y-coordinate (for click, mousemove events)
	Target          string    `json:"target,omitempty"`            // Target element (for click, hover events)
	ScrollPercentage int      `json:"scrollPercentage,omitempty"`  // Scroll depth as a percentage (for scroll events)
	Duration        int       `json:"duration,omitempty"`          // Duration in seconds (for duration, idle time events)
	ElementID       string    `json:"elementId,omitempty"`         // ID of the element (for hover, form, field events)
	ClassName       string    `json:"className,omitempty"`         // Class name of the element (for hover, form events)
	FieldName       string    `json:"fieldName,omitempty"`         // Name of the form field (for field focus/blur events)
	FormID          string    `json:"formId,omitempty"`            // ID of the form (for form submission events)
	FormClassName   string    `json:"formClassName,omitempty"`     // Class name of the form (for form submission events)
	Language        string    `json:"language,omitempty"`          // User's preferred language setting
	TimeZone        string    `json:"timeZone,omitempty"`          // User's time zone
	ScreenResolution string   `json:"screenResolution,omitempty"`  // Screen resolution of the user's device
	VideoID         string    `json:"videoId,omitempty"`           // ID of the video element (for video play/completion events)
	VideoURL        string    `json:"videoUrl,omitempty"`          // URL of the video (for video play/completion events)
	AudioID         string    `json:"audioId,omitempty"`           // ID of the audio element (for audio play events)
	AudioURL        string    `json:"audioUrl,omitempty"`          // URL of the audio (for audio play events)
	DownloadURL     string    `json:"downloadUrl,omitempty"`       // URL of the downloaded file (for download events)
	ImageURL        string    `json:"imageUrl,omitempty"`          // URL of the viewed image (for image view events)
}
