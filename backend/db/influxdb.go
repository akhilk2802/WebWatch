package db

import (
	"backend/config"
	"context"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var client influxdb2.Client
var writeAPI api.WriteAPIBlocking

func InitInfluxDB() {
	client = influxdb2.NewClient("http://localhost:8086", config.AppConf.InfluxToken)
	writeAPI = client.WriteAPIBlocking(config.AppConf.InfluxOrganisation, config.AppConf.InfluxBucket)
}

func writeData(p *write.Point) {
	log.Printf("Writing data to InfluxDB : %v", p)
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Fatalf("Failed to write data to InfluxDB: %v", err)
	}
}

func createPoint(measurement string, tags map[string]string, fields map[string]interface{}) *write.Point {
	return influxdb2.NewPoint(measurement, tags, fields, time.Now())
}

func storeData(measurement string, tags map[string]string, fields map[string]interface{}) {
	p := createPoint(measurement, tags, fields)
	writeData(p)
}

func StorePageViewData(url string, count int) {
	storeData("pageviews", map[string]string{"url": url, "data_type": "raw"}, map[string]interface{}{"count": count})
}

func StoreClickData(url string, target string, count int) {
	storeData("clicks", map[string]string{"url": url, "target": target, "data_type": "raw"}, map[string]interface{}{"count": count})
}

func StoreSessionDurationData(url string, avgDuration int) {
	storeData("session_durations", map[string]string{"url": url, "data_type": "raw"}, map[string]interface{}{"avg_duration": avgDuration})
}

func StoreScrollData(url string, scrollPercentage int) {
	storeData("scrolls", map[string]string{"url": url, "data_type": "raw"}, map[string]interface{}{"scroll_percentage": scrollPercentage})
}

func StoreMouseMovementData(url string, x int, y int) {
	storeData("mouse_movements", map[string]string{"url": url, "data_type": "raw"}, map[string]interface{}{"x": x, "y": y})
}

func StoreHoverData(url string, target string, id string, className string) {
	storeData("hovers", map[string]string{"url": url, "target": target, "id": id, "className": className, "data_type": "raw"}, map[string]interface{}{"hover_count": 1})
}

func StoreFormSubmissionData(formID string, formClassName string, url string) {
	storeData("form_submissions", map[string]string{"form_id": formID, "form_className": formClassName, "url": url, "data_type": "raw"}, map[string]interface{}{"submission_count": 1})
}

func StoreFieldFocusData(fieldID string, fieldName string, url string) {
	storeData("field_focuses", map[string]string{"field_id": fieldID, "field_name": fieldName, "url": url, "data_type": "raw"}, map[string]interface{}{"focus_count": 1})
}

func StoreFieldBlurData(fieldID string, fieldName string, url string) {
	storeData("field_blurs", map[string]string{"field_id": fieldID, "field_name": fieldName, "url": url, "data_type": "raw"}, map[string]interface{}{"blur_count": 1})
}

func StoreIdleTimeData(url string, duration int) {
	storeData("idle_times", map[string]string{"url": url, "data_type": "raw"}, map[string]interface{}{"duration": duration})
}

func StoreVideoPlayData(url string, videoID string, videoURL string) {
	storeData("video_plays", map[string]string{"url": url, "video_id": videoID, "video_url": videoURL, "data_type": "raw"}, map[string]interface{}{"play_count": 1})
}

func StoreVideoCompletionData(url string, videoID string, videoURL string) {
	storeData("video_completions", map[string]string{"url": url, "video_id": videoID, "video_url": videoURL, "data_type": "raw"}, map[string]interface{}{"completion_count": 1})
}

func StoreAudioPlayData(url string, audioID string, audioURL string) {
	storeData("audio_plays", map[string]string{"url": url, "audio_id": audioID, "audio_url": audioURL, "data_type": "raw"}, map[string]interface{}{"play_count": 1})
}

func StoreDownloadData(url string, downloadURL string) {
	storeData("downloads", map[string]string{"url": url, "download_url": downloadURL, "data_type": "raw"}, map[string]interface{}{"download_count": 1})
}

func StoreImageViewData(url string, imageURL string) {
	storeData("image_views", map[string]string{"url": url, "image_url": imageURL, "data_type": "raw"}, map[string]interface{}{"view_count": 1})
}

// Aggregation functions

func AggregatePageViewData(url string, count int) {
	storeData("pageviews", map[string]string{"url": url, "data_type": "aggregated"}, map[string]interface{}{"count": count})
}

func AggregateClickData(url, target string, count int) {
	storeData("clicks", map[string]string{"url": url, "target": target, "data_type": "aggregated"}, map[string]interface{}{"count": count})
}

func AggregateSessionDurationData(url string, avgDuration int) {
	storeData("session_durations", map[string]string{"url": url, "data_type": "aggregated"}, map[string]interface{}{"avg_duration": avgDuration})
}

func AggregateScrollData(url string, avgScrollPercentage int) {
	storeData("scrolls", map[string]string{"url": url, "data_type": "aggregated"}, map[string]interface{}{"avg_scroll_percentage": avgScrollPercentage})
}

func AggregateMouseMovementData(url string, avgX int, avgY int) {
	storeData("mouse_movements", map[string]string{"url": url, "data_type": "aggregated"}, map[string]interface{}{"avg_x": avgX, "avg_y": avgY})
}

func AggregateHoverData(url string, target string, hoverCount int) {
	storeData("hovers", map[string]string{"url": url, "target": target, "data_type": "aggregated"}, map[string]interface{}{"hover_count": hoverCount})
}

func AggregateFormSubmissionData(formID string, formClassName string, submissionCount int) {
	storeData("form_submissions", map[string]string{"form_id": formID, "form_className": formClassName, "data_type": "aggregated"}, map[string]interface{}{"submission_count": submissionCount})
}

func AggregateFieldFocusData(fieldID string, fieldName string, focusCount int) {
	storeData("field_focuses", map[string]string{"field_id": fieldID, "field_name": fieldName, "data_type": "aggregated"}, map[string]interface{}{"focus_count": focusCount})
}

func AggregateFieldBlurData(fieldID string, fieldName string, blurCount int) {
	storeData("field_blurs", map[string]string{"field_id": fieldID, "field_name": fieldName, "data_type": "aggregated"}, map[string]interface{}{"blur_count": blurCount})
}

func AggregateIdleTimeData(url string, avgDuration int) {
	storeData("idle_times", map[string]string{"url": url, "data_type": "aggregated"}, map[string]interface{}{"avg_duration": avgDuration})
}

func AggregateVideoPlayData(url string, videoID string, playCount int) {
	storeData("video_plays", map[string]string{"url": url, "video_id": videoID, "data_type": "aggregated"}, map[string]interface{}{"play_count": playCount})
}

func AggregateVideoCompletionData(url string, videoID string, completionCount int) {
	storeData("video_completions", map[string]string{"url": url, "video_id": videoID, "data_type": "aggregated"}, map[string]interface{}{"completion_count": completionCount})
}

func AggregateAudioPlayData(url string, audioID string, playCount int) {
	storeData("audio_plays", map[string]string{"url": url, "audio_id": audioID, "data_type": "aggregated"}, map[string]interface{}{"play_count": playCount})
}

func AggregateDownloadData(url string, downloadURL string, downloadCount int) {
	storeData("downloads", map[string]string{"url": url, "download_url": downloadURL, "data_type": "aggregated"}, map[string]interface{}{"download_count": downloadCount})
}

func AggregateImageViewData(url string, imageURL string, viewCount int) {
	storeData("image_views", map[string]string{"url": url, "image_url": imageURL, "data_type": "aggregated"}, map[string]interface{}{"view_count": viewCount})
}
