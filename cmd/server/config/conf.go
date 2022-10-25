package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Port     int      `yaml:"port"`
	RepoType RepoType `yaml:"repo_type"`
	MySQL    MySQL    `yaml:"mysql"`
	Dummy    Dummy    `yaml:"dummy"`
}

type RepoType string

const (
	RepoTypeMySQL RepoType = "mysql"
	RepoTypeDummy RepoType = "dummy"
)

type MySQL struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Dummy struct{}

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

func IsValidRepoType(t RepoType) bool {
	return t == RepoTypeMySQL || t == RepoTypeDummy
}
