package main

import (
	"io"
	"log"
	"net"

	"gorm.io/gorm"
)

func handleConn(conn *net.TCPConn, db *gorm.DB) {
	defer handleConnClose(conn)

	for {
		cmd, err := getRedisClientCmd(conn)
		if err != nil {
			return
		}

		log.Printf("received command \"%s\" from %s", cmd.ToString(), cmd.IP)
		if db.Create(toMySQLRecord(cmd)).Error != nil {
			return
		}

		res := handleRedisCmd(cmd)
		if _, err := io.WriteString(conn, res); err != nil {
			return
		}
	}
}

func handleConnClose(conn *net.TCPConn) {
	conn.Close()
	log.Printf("connection from %s closed\n", conn.RemoteAddr().String())
}
