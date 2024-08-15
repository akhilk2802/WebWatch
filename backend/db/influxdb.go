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
	// Replace with your InfluxDB connection details
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

func StorePageViewData(url string, count int) {
	p := influxdb2.NewPoint("pageviews",
		map[string]string{
			"url":       url,
			"data_type": "raw",
		},
		map[string]interface{}{
			"count": count,
		},
		time.Now())
	writeData(p)
}
func StoreClickData(url string, target string, count int) {
	p := influxdb2.NewPoint("clicks",
		map[string]string{"url": url, "target": target, "data_type": "raw"},
		map[string]interface{}{
			"count": count,
		},
		time.Now())
	writeData(p)
}
func StoreSessionDurationData(url string, avgDuration int) {
	p := influxdb2.NewPoint("session_durations",
		map[string]string{"url": url, "data_type": "raw"},
		map[string]interface{}{"avg_duration": avgDuration},
		time.Now())
	writeData(p)
}

func AggregatePageViewData(url string, count int) {
	log.Println("adding pageview to DB")
	p := influxdb2.NewPoint("pageviews",
		map[string]string{
			"url":       url,
			"data_type": "aggregated",
		},
		map[string]interface{}{"count": count},
		time.Now())
	writeData(p)
}

func AggregateClickData(url, target string, count int) {
	p := influxdb2.NewPoint("clicks",
		map[string]string{"url": url, "target": target, "data_type": "aggregated"},
		map[string]interface{}{"count": count},
		time.Now())
	writeData(p)
}

func AggregateSessionDurationData(url string, avgDuration int) {
	p := influxdb2.NewPoint("session_durations",
		map[string]string{"url": url, "data_type": "aggregated"},
		map[string]interface{}{"avg_duration": avgDuration},
		time.Now())
	writeData(p)
}
