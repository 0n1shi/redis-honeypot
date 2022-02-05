package main

import (
	"log"
	"net"
)

func startServer(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("connection establised from %s\n", conn.RemoteAddr().String())

		go handleConn(conn)
	}
}
