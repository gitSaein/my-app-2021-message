package config

import (
	"io/ioutil"
	"log"
	filePath "my-app-2021-message/constants"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Type              string `yaml:"Type"`
	MessageDatabase   string `yaml:"MessageDatabase"`
	MessageCollection string `yaml:"MessageCollection"`
	Uri               string `yaml:"Uri"`
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
