package main

import (
	"log"
	"net"
	"os"
)

const (
	proto = "tcp"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6379"
	}

	tcpAddr, err := net.ResolveTCPAddr(proto, ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP(proto, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("starting Beehive Redis server ...")
	startServer(listener)
}
