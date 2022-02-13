package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const LOCAL_PROPERTIEIES = "../config/properties-local.yaml"

type MongoDB struct {
	Database              string `yaml:"database"`
	MessageCollection     string `yaml:"messageCollection"`
	ParticipantCollection string `yaml:"participantCollection"`
	ChatCollection        string `yaml:"chatCollection"`
	Uri                   string `yaml:"uri"`
	User                  string `yaml:"user"`
	Pwd                   string `yaml:"pwd"`
}

type RabbitMQ struct {
	Url               string `yaml:"url"`
	QueueNameByUser   string `yaml:"queueNameByUser"`
	ExchangeName      string `yaml:"exchangeName"`
	MessageRoutingKey string `yaml:"messageRoutingKey"`
}

type DatabaseConfig struct {
	MongoDB  MongoDB  `yaml:"mongodb"`
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
}

// Config Type
type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func propertiesByEnvironment(env string) string {
	switch env {
	case "local":
		return LOCAL_PROPERTIEIES
	default:
		return ""
	}
}

func GetCongif(env string) Config {
	path := propertiesByEnvironment(env)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error of read file: %v", err)
	}

	configMap := &Config{}
	err = yaml.Unmarshal(buf, configMap)
	if err != nil {
		log.Fatalf("error of yaml unmarshal: %v", err)
	}
	return *configMap
}
