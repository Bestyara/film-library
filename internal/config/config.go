package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port       string `yaml:"port"`
		SecurePort string `yaml:"secure_port"`
	}
	Database struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		User   string `yaml:"user"`
		Dbname string `yaml:"dbname"`
	}
	Users []struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password,omitempty"`
	}
}

func ConfigInit() (Config, error) {
	var y Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println("Failed to read config.yaml")
	}
	err = yaml.Unmarshal(data, &y)
	if err != nil {
		log.Println("Failed to read config.yaml")
	}
	return y, err
}
