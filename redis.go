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
	IP     string
}

func (c *RedisCommand) ToString() string {
	str := c.Cmd
	if len(c.Args) > 0 {
		str += " " + strings.Join(c.Args, " ")
	}
	return str
}

func getRedisClientCmd(conn *net.TCPConn) (*RedisCommand, error) {
	buffer, err := readTCPPayload(conn)
	if err != nil {
		return nil, err
	}
	strs := parseRedisRawClientCmd2Strs(buffer)
	cmd, err := parseRedisClientCmd(strs)
	if err != nil {
		return nil, err
	}

	cmd.IP = strings.Split(conn.RemoteAddr().String(), ":")[0]
	return cmd, nil
}

func handleRedisCmd(cmd *RedisCommand) string {
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
	case "INFO":
		return redisINFO()
	case "CONFIG":
		return redisCONFIG()
	}
	return fmt.Sprintf("-ERR unknown command `%s`, with args beginning with: %s", cmd.Cmd, strings.Join(cmd.Args, " "))
}
