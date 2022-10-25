package redishoneypot

import (
	"log"
	"net"
)

const (
	proto = "tcp"
)

func StartServer(host string, repo Repository) {
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

		go handleConn(conn, repo)
	}
}
