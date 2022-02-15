package main

import (
	"fmt"
	"net"
	"strings"
)

type RedisCmdType string

const (
	RedisCmdCOMMAND RedisCmdType = "COMMAND"
	RedisCmdPING    RedisCmdType = "PING"
	RedisCmdKEYS    RedisCmdType = "KEYS"
	RedisCmdSET     RedisCmdType = "SET"
	RedisCmdGET     RedisCmdType = "GET"
	RedisCmdDEL     RedisCmdType = "DEL"
	RedisCmdINFO    RedisCmdType = "INFO"
	RedisCmdCONFIG  RedisCmdType = "CONFIG"
	RedisCmdSAVE    RedisCmdType = "SAVE"
)

var implemenetedCMDs = []RedisCmdType{
	RedisCmdCOMMAND,
	RedisCmdPING,
	RedisCmdKEYS,
	RedisCmdSET,
	RedisCmdGET,
	RedisCmdDEL,
	RedisCmdINFO,
	RedisCmdCONFIG,
}

func IsImplemented(cmd RedisCmdType) bool {
	for _, c := range implemenetedCMDs {
		if c == cmd {
			return true
		}
	}
	return false
}

type RedisCommand struct {
	Length      int
	Cmd         RedisCmdType
	Args        []string
	IP          string
	Implemented bool
}

func (c *RedisCommand) ToString() string {
	str := string(c.Cmd)
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
	cmd.Implemented = IsImplemented(cmd.Cmd)
	return cmd, nil
}

func handleRedisCmd(cmd *RedisCommand) string {
	switch cmd.Cmd {
	case RedisCmdCOMMAND:
		return redisCOMMAND()
	case RedisCmdPING:
		return redisPING()
	case RedisCmdKEYS:
		return redisKEYS()
	case RedisCmdSET:
		return redisSET(cmd.Args)
	case RedisCmdGET:
		return redisGET(cmd.Args[0])
	case RedisCmdDEL:
		return redisDEL(cmd.Args[0])
	case RedisCmdINFO:
		return redisINFO()
	case RedisCmdCONFIG:
		return redisCONFIG()
	case RedisCmdSAVE:
		return redisSAVE()
	}
	return fmt.Sprintf("-ERR unknown command `%s`, with args beginning with: %s", cmd.Cmd, strings.Join(cmd.Args, " "))
}
