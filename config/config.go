package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version     string `yaml:"version"`
	Blockcypher struct {
		Token string `yaml:"token"`
	} `yaml:"blockcypher"`
}

// ParseConfig from config.yml
func ParseConfig() Config {
	c := Config{}

	data, _err := ioutil.ReadFile("config.yml")
	if _err != nil {
		log.Printf("config.Get err   #%v ", _err)
	}

	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return c
}
