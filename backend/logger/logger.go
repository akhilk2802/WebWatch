package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

// InitLogger initializes the logger and sets the output destination
func InitLogger() {
	// Create or open the log file
	file, err := os.OpenFile("webwatch.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Initialize the logger
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(v ...interface{}) {
	Logger.Println(v...)
}

// Fatal logs fatal messages and exits the application
func Fatal(v ...interface{}) {
	Logger.Fatal(v...)
}
