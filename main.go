package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	var confName string
	flag.StringVar(&confName, "conf", "", "config file path")
	flag.Parse()

	conf, err := getConf(confName)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initMySQL(&conf.MySQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting Beehive Redis server ...")
	startServer(fmt.Sprintf(":%d", conf.Redis.Port), db)
}

func getConf(filename string) (*Conf, error) {
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
