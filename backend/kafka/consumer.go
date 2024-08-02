package kafka

import (
	"backend/config"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(topic string) {
	// log.Printf("Here is the topic : %s", topic)
	// Create a new consumer group
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.AppConf.KafkaBrokerURL},
		Topic:    topic,
		GroupID:  config.AppConf.KafkaGroupId,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("could not read message: %v", err)
			// Optionally, add a retry mechanism or log the error for monitoring purposes
			continue
		}
		log.Printf("message at offset %d: key=%s value=%s\n", m.Offset, string(m.Key), string(m.Value))
		// Process the message here
	}

}

func StartAllConsumers() {
	topics := []string{
		config.AppConf.KafkaTopicPageView,
		config.AppConf.KafkaTopicClick,
		config.AppConf.KafkaTopicDuration,
	}

	for _, topic := range topics {
		go StartConsumer(topic)
	}
}
