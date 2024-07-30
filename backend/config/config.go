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
	AppConf.KafkaTopicPageView = viper.GetString("KAFKA_TOPIC_PAGE_VIEW")
	AppConf.KafkaTopicClick = viper.GetString("KAFKA_TOPIC_CLICK")
	AppConf.KafkaTopicDuration = viper.GetString("KAFKA_TOPIC_DURATION")
	AppConf.ServerPort = viper.GetString("SERVER_PORT")

}
