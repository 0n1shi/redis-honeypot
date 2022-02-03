package main

import (
	"bufio"
	"io"
)

func read1k(rd io.Reader) ([]byte, error) {
	buffer := make([]byte, bufferSize)
	_, err := bufio.NewReader(rd).Read(buffer)
	return buffer, err
}
