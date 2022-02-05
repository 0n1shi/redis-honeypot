package main

import (
	"io"
	"log"
	"net"
	"strings"
)

func handleConn(conn *net.TCPConn) {
	defer handleConnClose(conn)

	for {
		cmd, err := getRedisClientCmd(conn)
		if err != nil {
			return
		}

		cmdStr := cmd.Cmd
		if len(cmd.Args) > 0 {
			cmdStr += " " + strings.Join(cmd.Args, " ")
		}
		log.Printf("received command \"%s\" from %s", cmdStr, conn.RemoteAddr().String())

		res := handleRedisCommand(cmd)
		if _, err := io.WriteString(conn, res); err != nil {
			return
		}
	}
}

func handleConnClose(conn *net.TCPConn) {
	conn.Close()
	log.Printf("connection from %s closed\n", conn.RemoteAddr().String())
}
