package config

import (
	"io/ioutil"
	"log"
	filePath "my-app-2021-message/constants"

	"gopkg.in/yaml.v2"
)

type MongoDB struct {
	MessageDatabase   string `yaml:"messageDatabase"`
	MessageCollection string `yaml:"messageCollection"`
	Uri               string `yaml:"uri"`
	User              string `yaml:"user"`
	Pwd               string `yaml:"pwd"`
}

type RabbitMQ struct {
	Url string `yaml:"url"`
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
		return filePath.LOCAL_PROPERTIEIES
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
