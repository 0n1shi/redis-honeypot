package main

import (
	"io"
	"log"
	"net"
	"strings"
)

const bufferSize = 1024

func comunicate(conn *net.TCPConn) {
	defer conn.Close()
	defer handleConnClose(conn)

	for {
		buffer, err := read1k(conn)
		if err != nil {
			return
		}

		rawCMDStrs := cmdStrs(buffer)
		clientCMD, err := parseCmd(rawCMDStrs)
		if err != nil {
			log.Printf("failed to parse \"%+v\"", rawCMDStrs)
			return
		}

		cmdStr := clientCMD.Cmd
		if len(clientCMD.Args) > 0 {
			cmdStr += " " + strings.Join(clientCMD.Args, " ")
		}
		log.Printf("received command \"%s\" from %s", cmdStr, conn.RemoteAddr().String())

		if _, err := io.WriteString(conn, "+PONG\r\n"); err != nil {
			return
		}
	}
}

func handleConnClose(conn *net.TCPConn) {
	log.Printf("connection from %s closed\n", conn.RemoteAddr().String())
}
