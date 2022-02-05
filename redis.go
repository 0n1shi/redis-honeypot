package main

import (
	"fmt"
	"net"
	"strings"
)

type RedisCommand struct {
	Length int
	Cmd    string
	Args   []string
}

func getRedisClientCmd(conn *net.TCPConn) (*RedisCommand, error) {
	buffer, err := readTCPPayload(conn)
	if err != nil {
		return nil, err
	}
	strs := parseRedisRawClientCmd2Strs(buffer)
	return parseRedisClientCmd(strs)
}

func handleRedisCommand(cmd *RedisCommand) string {
	switch cmd.Cmd {
	case "COMMAND":
		return redisCOMMAND()
	case "PING":
		return redisPING()
	case "KEYS":
		return redisKEYS()
	case "SET":
		return redisSET(cmd.Args)
	case "GET":
		return redisGET(cmd.Args[0])
	case "DEL":
		return redisDEL(cmd.Args[0])
	}
	return fmt.Sprintf("-ERR unknown command `%s`, with args beginning with: %s", cmd.Cmd, strings.Join(cmd.Args, " "))
}
