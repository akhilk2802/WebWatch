package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	KafkaBrokerURL     string
	KafkaTopicPageView string
	KafkaTopicClick    string
	KafkaTopicDuration string
	ServerPort         string
	KafkaGroupId       string
	InfluxToken        string
	InfluxBucket       string
	InfluxOrganisation string
}

var AppConf Config

func InitConfig() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("/backend/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error While reading config file %s", err)
	}

	AppConf.KafkaBrokerURL = viper.GetString("KAFKA_BROKER_URL")
	AppConf.KafkaTopicPageView = viper.GetString("KAFKA_TOPIC_PAGEVIEW")
	AppConf.KafkaTopicClick = viper.GetString("KAFKA_TOPIC_CLICK")
	AppConf.KafkaTopicDuration = viper.GetString("KAFKA_TOPIC_DURATION")
	AppConf.ServerPort = viper.GetString("SERVER_PORT")
	AppConf.KafkaGroupId = viper.GetString("KAFKA_GROUP_ID")
	AppConf.InfluxToken = viper.GetString("INFLUX_TOKEN")
	AppConf.InfluxBucket = viper.GetString("INFLUX_BUCKET")
	AppConf.InfluxOrganisation = viper.GetString("INFLUX_ORGANISATION")

}
