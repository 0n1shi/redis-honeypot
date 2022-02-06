package main

import (
	"log"
	"net"

	"gorm.io/gorm"
)

const (
	proto = "tcp"
)

func startServer(host string, db *gorm.DB) {
	tcpAddr, err := net.ResolveTCPAddr(proto, host)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP(proto, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("connection establised from %s\n", conn.RemoteAddr().String())

		go handleConn(conn, db)
	}
}
