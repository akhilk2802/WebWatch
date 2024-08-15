package kafka

import (
	"backend/config"
	"backend/db"
	"backend/models"
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

var (
	pageViewCounts  = make(map[string]int)
	clickCounts     = make(map[string]map[string]int)
	sessionDuration = make(map[string][]int)
	mu              sync.Mutex
)

func StartConsumer(topic string) {
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
		}
		var event models.Event
		err = json.Unmarshal(m.Value, &event)
		if err != nil {
			log.Printf("could not unmarshal message: %v", err)
			continue
		}
		processEvent(event)
	}

}

func processEvent(event models.Event) {
	mu.Lock()
	defer mu.Unlock()
	switch event.Type {
	case "pageview":
		log.Printf("url : %s", event.URL)
		db.StorePageViewData(event.URL, 1)
		pageViewCounts[event.URL]++
	case "click":
		db.StoreClickData(event.URL, event.Target, 1)
		if clickCounts[event.URL] == nil {
			clickCounts[event.URL] = make(map[string]int)
		}
		clickCounts[event.URL][event.Target]++
	case "duration":
		db.StoreSessionDurationData(event.URL, event.Duration)
		sessionDuration[event.URL] = append(sessionDuration[event.URL], event.Duration)
	}

}

func AggregateData() {
	mu.Lock()
	defer mu.Unlock()

	log.Println("Aggregated Page Views: ")
	for url, count := range pageViewCounts {
		db.AggregatePageViewData(url, count)
		log.Printf("URL: %s, Page Views: %d\n", url, count)
	}
	log.Println("Aggregated Clicks: ")
	for url, clicks := range clickCounts {
		for target, count := range clicks {
			db.AggregateClickData(url, target, count)
			log.Printf("URL: %s, click id: %s, Clicks: %d\n", url, target, count)
		}
	}
	log.Println("Aggregated Session Durations: ")
	for url, durations := range sessionDuration {
		var totalDuration int
		for _, duration := range durations {
			totalDuration += duration
		}
		avgDuration := totalDuration / len(durations)
		db.AggregateSessionDurationData(url, avgDuration)
		log.Printf("URL: %s, Average session duration: %d seconds\n", url, avgDuration)
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
