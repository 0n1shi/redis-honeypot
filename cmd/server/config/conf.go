package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Port  int   `yaml:"port"`
	MySQL MySQL `yaml:"mysql"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

func GetContent(filename string) (*Conf, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var conf Conf
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatal(err)
	}

	return &conf, nil
}
