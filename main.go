package main

import (
	"bufio"
	"io"
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
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("connection establised from %s\n", conn.RemoteAddr().String())

		go (func() {
			defer conn.Close()
			defer (func() {
				log.Printf("connection from %s closed\n", conn.RemoteAddr().String())
			})()
			for {
				buffer := make([]byte, 1024)
				_, err := bufio.NewReader(conn).Read(buffer)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("%s\n", string(buffer))
				if _, err := io.WriteString(conn, "+PONG\r\n"); err != nil {
					return
				}
			}
		})()
	}
}
