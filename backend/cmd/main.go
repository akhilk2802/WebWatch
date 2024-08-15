package main

import (
	"backend/config"
	"backend/db"
	"backend/kafka"
	"backend/router"
	"log"
	"net/http"
	"time"
)

func main() {
	config.InitConfig()
	kafka.InitKafka()
	// db.InitMongoDB()
	db.InitInfluxDB()

	kafka.StartAllConsumers()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				kafka.AggregateData()
			}
		}
	}()

	r := router.InitRouter()

	log.Println("Starting server on port :", config.AppConf.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConf.ServerPort, r); err != nil {
		log.Fatalf("could not start server : %s\n", err.Error())
	}
}
