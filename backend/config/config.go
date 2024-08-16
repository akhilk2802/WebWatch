package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	KafkaBrokerURL string

	KafkaTopicPageView        string
	KafkaTopicClick           string
	KafkaTopicDuration        string
	KafkaTopicScroll          string
	KafkaTopicMouseMove       string
	KafkaTopicHover           string
	KafkaTopicFormSubmission  string
	KafkaTopicFieldFocus      string
	KafkaTopicFieldBlur       string
	KafkaTopicIdleTime        string
	KafkaTopicVideoPlay       string
	KafkaTopicVideoCompletion string
	KafkaTopicAudioPlay       string
	KafkaTopicDownload        string
	KafkaTopicImageView       string

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
	AppConf.KafkaTopicScroll = viper.GetString("KAFKA_TOPIC_SCROLL")
	AppConf.KafkaTopicMouseMove = viper.GetString("KAFKA_TOPIC_MOUSEMOVE")
	AppConf.KafkaTopicHover = viper.GetString("KAFKA_TOPIC_HOVER")
	AppConf.KafkaTopicFormSubmission = viper.GetString("KAFKA_TOPIC_FORMSUBMISSION")
	AppConf.KafkaTopicFieldFocus = viper.GetString("KAFKA_TOPIC_FIELDFOCUS")
	AppConf.KafkaTopicFieldBlur = viper.GetString("KAFKA_TOPIC_FIELDBLUR")
	AppConf.KafkaTopicIdleTime = viper.GetString("KAFKA_TOPIC_IDLETIME")
	AppConf.KafkaTopicVideoPlay = viper.GetString("KAFKA_TOPIC_VIDEOPLAY")
	AppConf.KafkaTopicVideoCompletion = viper.GetString("KAFKA_TOPIC_VIDEOCOMPLETION")
	AppConf.KafkaTopicAudioPlay = viper.GetString("KAFKA_TOPIC_AUDIOPLAY")
	AppConf.KafkaTopicDownload = viper.GetString("KAFKA_TOPIC_DOWNLOAD")
	AppConf.KafkaTopicImageView = viper.GetString("KAFKA_TOPIC_IMAGEVIEW")

	AppConf.ServerPort = viper.GetString("SERVER_PORT")
	AppConf.KafkaGroupId = viper.GetString("KAFKA_GROUP_ID")
	AppConf.InfluxToken = viper.GetString("INFLUX_TOKEN")
	AppConf.InfluxBucket = viper.GetString("INFLUX_BUCKET")
	AppConf.InfluxOrganisation = viper.GetString("INFLUX_ORGANISATION")

}
