package router

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.CORS)
	router.HandleFunc("/track", handlers.TrackEvent).Methods("POST", "OPTIONS")
	return router
}
