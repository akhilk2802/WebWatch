package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitMongoDB initializes the MongoDB client
func InitMongoDB() {
	log.Println("I am here")
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}

// StoreAggregatedData stores the aggregated data in MongoDB
func StoreAggregatedData(collectionName string, data interface{}) {
	collection := client.Database("analytics").Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatalf("Failed to store data: %v", err)
	}
}

// AggregatePageViewData stores aggregated page view data
func AggregatePageViewData(url string, count int) {
	data := bson.M{"url": url, "page_views": count, "timestamp": time.Now()}
	StoreAggregatedData("pageviews", data)
}

// AggregateClickData stores aggregated click data
func AggregateClickData(url, target string, count int) {
	data := bson.M{"url": url, "element_id": target, "clicks": count, "timestamp": time.Now()}
	StoreAggregatedData("clicks", data)
}

// AggregateSessionDurationData stores aggregated session duration data
func AggregateSessionDurationData(url string, avgDuration int) {
	data := bson.M{"url": url, "average_duration": avgDuration, "timestamp": time.Now()}
	StoreAggregatedData("durations", data)
}
