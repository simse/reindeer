package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var config struct {
	Meta struct {
		Title string `yaml:"title"`
	} `yaml:"meta"`
}

func ConfigPath() string {
	return "/home/simon/projects/reindeer/config.example.yaml"
}

func LoadConfig() {
	// read file
	configFile, err := ioutil.ReadFile(ConfigPath())

	if err != nil {
		log.Fatal(err)
	}

	// decode file
	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		log.Fatal(err)
	}
}

func GetString() {
	fmt.Println(config.Meta.Title)
}
