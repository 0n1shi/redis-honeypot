package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	honeypot "github.com/0n1shi/redis-honeypot"
	"github.com/0n1shi/redis-honeypot/cmd/server/config"
	"github.com/0n1shi/redis-honeypot/repository/dummy"
	"github.com/0n1shi/redis-honeypot/repository/mysql"
)

var (
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {
	app := &cli.App{
		Name:  "Redis server of Beehive honeypot series",
		Usage: "TCP server which communicate in Redis protocol.",
		Commands: []cli.Command{
			{
				Name:   "version",
				Usage:  "Show version",
				Action: showVersion,
			},
			{
				Name:  "run",
				Usage: "Start tcp server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "config",
						Value:    "",
						Required: true,
						Usage:    "config file path",
					},
				},
				Action: runServer,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func showVersion(c *cli.Context) error {
	fmt.Printf("version: %s\n", version)
	return nil
}

func runServer(c *cli.Context) error {
	confPath := c.String("config")
	if confPath == "" {
		log.Fatalln("expected config file path")
	}

	conf, err := config.GetContent(confPath)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	if !config.IsValidRepoType(conf.RepoType) {
		log.Fatalf("Invalid repository type: %s\n", conf.RepoType)
	}

	var repo honeypot.Repository
	if conf.RepoType == config.RepoTypeMySQL {
		repo, err = mysql.NewMySQLRepository(&mysql.Conf{
			Host:     conf.MySQL.Host,
			User:     conf.MySQL.User,
			Password: conf.MySQL.Password,
			DB:       conf.MySQL.DB,
		})
	}
	if conf.RepoType == config.RepoTypeDummy {
		repo = dummy.NewDummyRepository()
	}
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}

	log.Println("starting Beehive Redis server ...")
	honeypot.StartServer(fmt.Sprintf(":%d", conf.Port), repo)
	return nil
}
