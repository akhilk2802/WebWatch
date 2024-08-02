package kafka

import (
	"backend/config"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var Writer map[string]*kafka.Writer

func InitKafka() {
	Writer = make(map[string]*kafka.Writer)

	topicConfigs := []struct {
		name              string
		partitions        int
		replicationFactor int
	}{
		{"pageview-topic", 1, 1},
		{"click-topic", 1, 1},
		{"duration-topic", 1, 1},
	}

	for _, topicConfig := range topicConfigs {
		err := createTopic(topicConfig.name, topicConfig.partitions, topicConfig.replicationFactor)
		if err != nil {
			log.Fatalf("failed to create topic: %v", err)
		}
		Writer[topicConfig.name] = &kafka.Writer{
			Addr:     kafka.TCP(config.AppConf.KafkaBrokerURL),
			Topic:    topicConfig.name,
			Balancer: &kafka.LeastBytes{},
		}
	}
}

func createTopic(topic string, partitions int, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", config.AppConf.KafkaBrokerURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	conn, err = kafka.Dial("tcp", controller.Host)
	if err != nil {
		return err
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}

func ProduceMessage(eventType string, key string, message []byte) {

	writer, exists := Writer[eventType]
	if !exists {
		log.Printf("No Writer configured for eventType: %s\n", eventType)

		// Try to create the topic if it doesn't exist
		err := createTopic(eventType, 1, 1)
		if err != nil {
			log.Fatalf("failed to create topic: %v", err)
			return
		}

		// Create a new writer for the new topic
		writer = &kafka.Writer{
			Addr:     kafka.TCP(config.AppConf.KafkaBrokerURL),
			Topic:    eventType,
			Balancer: &kafka.LeastBytes{},
		}
		Writer[eventType] = writer
	}

	log.Printf("Event Type : %s", eventType)
	log.Printf("Message : %s", message)

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
