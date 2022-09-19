package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Db Db `yaml:"db"`
}

func Load() (*Config, error) {

	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalln("Failed load config.", err)
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln("Failed close connection.", err)
		}
	}(f)

	var config Config
	err = yaml.NewDecoder(f).Decode(&config)

	return &config, err
}
