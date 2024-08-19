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
	pageViewCounts        = make(map[string]int)
	clickCounts           = make(map[string]map[string]int)
	sessionDuration       = make(map[string][]int)
	scrollData            = make(map[string][]int)
	mouseMovementData     = make(map[string][][2]int) // Store X and Y coordinates as pairs
	hoverCounts           = make(map[string]map[string]int)
	formSubmissionCounts  = make(map[string]int)
	fieldFocusCounts      = make(map[string]int)
	fieldBlurCounts       = make(map[string]int)
	idleTimes             = make(map[string][]int)
	videoPlayCounts       = make(map[string]map[string]int)
	videoCompletionCounts = make(map[string]map[string]int)
	audioPlayCounts       = make(map[string]map[string]int)
	downloadCounts        = make(map[string]int)
	imageViewCounts       = make(map[string]int)
	mu                    sync.Mutex
)

func StartConsumer(topic string) {

	if topic == "" {
		log.Fatal("Topic is empty")
	}

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

	case "scroll":
		db.StoreScrollData(event.URL, event.ScrollPercentage)
		scrollData[event.URL] = append(scrollData[event.URL], event.ScrollPercentage)

	case "mousemove":
		db.StoreMouseMovementData(event.URL, event.X, event.Y)
		mouseMovementData[event.URL] = append(mouseMovementData[event.URL], [2]int{event.X, event.Y})

	case "hover":
		db.StoreHoverData(event.URL, event.Target, event.ElementID, event.ClassName)
		if hoverCounts[event.URL] == nil {
			hoverCounts[event.URL] = make(map[string]int)
		}
		hoverCounts[event.URL][event.Target]++

	case "form_submission":
		db.StoreFormSubmissionData(event.FormID, event.FormClassName, event.URL)
		formSubmissionCounts[event.FormID]++

	case "field_focus":
		db.StoreFieldFocusData(event.ElementID, event.FieldName, event.URL)
		fieldFocusCounts[event.ElementID]++

	case "field_blur":
		db.StoreFieldBlurData(event.ElementID, event.FieldName, event.URL)
		fieldBlurCounts[event.ElementID]++

	case "idle_time":
		db.StoreIdleTimeData(event.URL, event.Duration)
		idleTimes[event.URL] = append(idleTimes[event.URL], event.Duration)

	case "video_play":
		db.StoreVideoPlayData(event.URL, event.VideoID, event.VideoURL)
		if videoPlayCounts[event.URL] == nil {
			videoPlayCounts[event.URL] = make(map[string]int)
		}
		videoPlayCounts[event.URL][event.VideoID]++

	case "video_completion":
		db.StoreVideoCompletionData(event.URL, event.VideoID, event.VideoURL)
		if videoCompletionCounts[event.URL] == nil {
			videoCompletionCounts[event.URL] = make(map[string]int)
		}
		videoCompletionCounts[event.URL][event.VideoID]++

	case "audio_play":
		db.StoreAudioPlayData(event.URL, event.AudioID, event.AudioURL)
		if audioPlayCounts[event.URL] == nil {
			audioPlayCounts[event.URL] = make(map[string]int)
		}
		audioPlayCounts[event.URL][event.AudioID]++

	case "download":
		db.StoreDownloadData(event.URL, event.DownloadURL)
		downloadCounts[event.URL]++

	case "image_view":
		db.StoreImageViewData(event.URL, event.ImageURL)
		imageViewCounts[event.URL]++

	default:
		log.Printf("Unhandled event type: %s", event.Type)
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

	log.Println("Aggregated Scroll Data: ")
	for url, scrolls := range scrollData {
		var totalScroll int
		for _, scroll := range scrolls {
			totalScroll += scroll
		}
		avgScroll := totalScroll / len(scrolls)
		db.AggregateScrollData(url, avgScroll)
		log.Printf("URL: %s, Average scroll percentage: %d\n", url, avgScroll)
	}

	log.Println("Aggregated Mouse Movements: ")
	for url, movements := range mouseMovementData {
		var totalX, totalY int
		for _, movement := range movements {
			totalX += movement[0]
			totalY += movement[1]
		}
		avgX := totalX / len(movements)
		avgY := totalY / len(movements)
		db.AggregateMouseMovementData(url, avgX, avgY)
		log.Printf("URL: %s, Average X: %d, Average Y: %d\n", url, avgX, avgY)
	}

	log.Println("Aggregated Hover Data: ")
	for url, hovers := range hoverCounts {
		for target, count := range hovers {
			db.AggregateHoverData(url, target, count)
			log.Printf("URL: %s, Target: %s, Hover Count: %d\n", url, target, count)
		}
	}

	log.Println("Aggregated Form Submissions: ")
	for formID, count := range formSubmissionCounts {
		db.AggregateFormSubmissionData(formID, "", count)
		log.Printf("Form ID: %s, Submission Count: %d\n", formID, count)
	}

	log.Println("Aggregated Field Focuses: ")
	for fieldID, count := range fieldFocusCounts {
		db.AggregateFieldFocusData(fieldID, "", count)
		log.Printf("Field ID: %s, Focus Count: %d\n", fieldID, count)
	}

	log.Println("Aggregated Field Blurs: ")
	for fieldID, count := range fieldBlurCounts {
		db.AggregateFieldBlurData(fieldID, "", count)
		log.Printf("Field ID: %s, Blur Count: %d\n", fieldID, count)
	}

	log.Println("Aggregated Idle Times: ")
	for url, times := range idleTimes {
		var totalDuration int
		for _, duration := range times {
			totalDuration += duration
		}
		avgDuration := totalDuration / len(times)
		db.AggregateIdleTimeData(url, avgDuration)
		log.Printf("URL: %s, Average Idle Duration: %d seconds\n", url, avgDuration)
	}

	log.Println("Aggregated Video Plays: ")
	for url, videoPlays := range videoPlayCounts {
		for videoID, count := range videoPlays {
			db.AggregateVideoPlayData(url, videoID, count)
			log.Printf("URL: %s, Video ID: %s, Play Count: %d\n", url, videoID, count)
		}
	}

	log.Println("Aggregated Video Completions: ")
	for url, videoCompletions := range videoCompletionCounts {
		for videoID, count := range videoCompletions {
			db.AggregateVideoCompletionData(url, videoID, count)
			log.Printf("URL: %s, Video ID: %s, Completion Count: %d\n", url, videoID, count)
		}
	}

	log.Println("Aggregated Audio Plays: ")
	for url, audioPlays := range audioPlayCounts {
		for audioID, count := range audioPlays {
			db.AggregateAudioPlayData(url, audioID, count)
			log.Printf("URL: %s, Audio ID: %s, Play Count: %d\n", url, audioID, count)
		}
	}

	log.Println("Aggregated Downloads: ")
	for url, count := range downloadCounts {
		db.AggregateDownloadData(url, url, count)
		log.Printf("URL: %s, Download Count: %d\n", url, count)
	}

	log.Println("Aggregated Image Views: ")
	for url, count := range imageViewCounts {
		db.AggregateImageViewData(url, url, count)
		log.Printf("URL: %s, Image View Count: %d\n", url, count)
	}
}

// func StartAllConsumers() {
// 	topics := []string{
// 		config.AppConf.KafkaTopicPageView,
// 		config.AppConf.KafkaTopicClick,
// 		config.AppConf.KafkaTopicDuration,
// 		config.AppConf.KafkaTopicScroll,
// 		config.AppConf.KafkaTopicMouseMove,
// 		config.AppConf.KafkaTopicHover,
// 		config.AppConf.KafkaTopicFormSubmission,
// 		config.AppConf.KafkaTopicFieldFocus,
// 		config.AppConf.KafkaTopicFieldBlur,
// 		config.AppConf.KafkaTopicIdleTime,
// 		config.AppConf.KafkaTopicVideoPlay,
// 		config.AppConf.KafkaTopicVideoCompletion,
// 		config.AppConf.KafkaTopicAudioPlay,
// 		config.AppConf.KafkaTopicDownload,
// 		config.AppConf.KafkaTopicImageView,
// 	}

// 	for _, topic := range topics {
// 		if topic != "" {
// 			go StartConsumer(topic)
// 		} else {
// 			log.Println("Skipping empty topic: ", topic)
// 		}
// 	}
// }

func StartAllConsumers() {
	topics := map[string]string{
		"KafkaTopicPageView":        config.AppConf.KafkaTopicPageView,
		"KafkaTopicClick":           config.AppConf.KafkaTopicClick,
		"KafkaTopicDuration":        config.AppConf.KafkaTopicDuration,
		"KafkaTopicScroll":          config.AppConf.KafkaTopicScroll,
		"KafkaTopicMouseMove":       config.AppConf.KafkaTopicMouseMove,
		"KafkaTopicHover":           config.AppConf.KafkaTopicHover,
		"KafkaTopicFormSubmission":  config.AppConf.KafkaTopicFormSubmission,
		"KafkaTopicFieldFocus":      config.AppConf.KafkaTopicFieldFocus,
		"KafkaTopicFieldBlur":       config.AppConf.KafkaTopicFieldBlur,
		"KafkaTopicIdleTime":        config.AppConf.KafkaTopicIdleTime,
		"KafkaTopicVideoPlay":       config.AppConf.KafkaTopicVideoPlay,
		"KafkaTopicVideoCompletion": config.AppConf.KafkaTopicVideoCompletion,
		"KafkaTopicAudioPlay":       config.AppConf.KafkaTopicAudioPlay,
		"KafkaTopicDownload":        config.AppConf.KafkaTopicDownload,
		"KafkaTopicImageView":       config.AppConf.KafkaTopicImageView,
	}

	for key, topic := range topics {
		if topic != "" {
			log.Printf("Starting consumer for topic: %s (config key: %s)\n", topic, key)
			go StartConsumer(topic)
		} else {
			log.Printf("Skipping empty topic: (config key: %s)\n", key)
		}
	}
}
