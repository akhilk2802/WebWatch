package main

import (
	"backend/config"
	"backend/kafka"
	"backend/router"
	"log"
	"net/http"
)

func main() {
	config.InitConfig()
	kafka.InitKafka()

	kafka.StartAllConsumers()

	r := router.InitRouter()

	log.Println("Starting server on port :", config.AppConf.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConf.ServerPort, r); err != nil {
		log.Fatalf("could not start server : %s\n", err.Error())
	}
}
