package main

import (
	"bufio"
	"net"
)

const bufferSize = 1024

func readTCPPayload(conn *net.TCPConn) ([]byte, error) {
	buffer := make([]byte, bufferSize)
	_, err := bufio.NewReader(conn).Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
