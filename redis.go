package main

import (
	"strconv"
	"strings"
)

type ClientCommand struct {
	Length int
	Cmd    string
	Args   []string
}

func cmdStrs(buffer []byte) []string {
	cmdStr := string(buffer)
	strs := strings.Split(cmdStr, "\r\n")
	strs = strs[:len(strs)-1]
	return strs
}

func parseCmd(strs []string) (*ClientCommand, error) {
	length, err := strconv.Atoi(strs[0][1:])
	if err != nil {
		return nil, err
	}
	cmd := ClientCommand{}
	if length <= 0 {
		return &cmd, nil
	}
	cmd.Length = length
	cmd.Cmd = strings.ToUpper(strs[2])
	for i := 3; i < len(strs); i = i + 2 {
		cmd.Args = append(cmd.Args, strs[i+1])
	}
	return &cmd, nil
}
