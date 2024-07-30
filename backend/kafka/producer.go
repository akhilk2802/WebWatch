package kafka

import (
	"backend/config"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// var Writer *kafka.Writer
var Writer map[string]*kafka.Writer

// func InitKafka() {

// 	Writer = &kafka.Writer{
// 		Addr:     kafka.TCP(config.AppConf.KafkaBrokerURL),
// 		Topic:    config.AppConf.KafkaTopic,
// 		Balancer: &kafka.LeastBytes{},
// 	}
// }

func InitKafka() {
	Writer = make(map[string]*kafka.Writer)
	Writer["pageview"] = &kafka.Writer{
		Addr:     kafka.TCP(config.AppConf.KafkaBrokerURL),
		Topic:    "pageview-topic",
		Balancer: &kafka.LeastBytes{},
	}
	Writer["click"] = &kafka.Writer{
		Addr:     kafka.TCP(config.AppConf.KafkaBrokerURL),
		Topic:    "click-topic",
		Balancer: &kafka.LeastBytes{},
	}
	Writer["duration"] = &kafka.Writer{
		Addr:     kafka.TCP(config.AppConf.KafkaBrokerURL),
		Topic:    "duration-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func ProduceMessage(eventType string, key string, message []byte) {

	writer, exists := Writer[eventType]
	if !exists {
		log.Printf("No Writer configured for eventType: %s\n", eventType)
		return
	}

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: message,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages: ", err)
	}
}
