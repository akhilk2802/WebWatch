package kafka

import (
	"backend/config"
	"backend/logger"
	"context"

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
		{"pageview", 1, 1},
		{"click", 1, 1},
		{"duration", 1, 1},
		{"scroll", 1, 1},
		{"mousemove", 1, 1},
		{"hover", 1, 1},
		{"form_submission", 1, 1},
		{"field_focus", 1, 1},
		{"field_blur", 1, 1},
		{"idle_time", 1, 1},
		{"video_play", 1, 1},
		{"video_completion", 1, 1},
		{"audio_play", 1, 1},
		{"download", 1, 1},
		{"image_view", 1, 1},
	}

	for _, topicConfig := range topicConfigs {
		err := createTopic(topicConfig.name, topicConfig.partitions, topicConfig.replicationFactor)
		if err != nil {
			logger.Logger.Fatalf("failed to create topic: %v", err)
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
		logger.Logger.Printf("No Writer configured for eventType: %s\n", eventType)

		// Try to create the topic if it doesn't exist
		err := createTopic(eventType, 1, 1)
		if err != nil {
			logger.Logger.Fatalf("failed to create topic: %v", err)
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

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: message,
		},
	)
	if err != nil {
		logger.Logger.Fatal("failed to write messages: ", err)
	}
}
