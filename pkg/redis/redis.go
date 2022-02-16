package redis

import (
	"fmt"
	"net"
	"strings"
)

type RedisCmd string

const (
	CmdCOMMAND RedisCmd = "COMMAND"
	CmdPING    RedisCmd = "PING"
	CmdKEYS    RedisCmd = "KEYS"
	CmdSET     RedisCmd = "SET"
	CmdGET     RedisCmd = "GET"
	CmdDEL     RedisCmd = "DEL"
	CmdINFO    RedisCmd = "INFO"
	CmdCONFIG  RedisCmd = "CONFIG"
	CmdSAVE    RedisCmd = "SAVE"
)

var implemenetedCMDs = []RedisCmd{
	CmdCOMMAND,
	CmdPING,
	CmdKEYS,
	CmdSET,
	CmdGET,
	CmdDEL,
	CmdINFO,
	CmdCONFIG,
}

func IsImplemented(cmd RedisCmd) bool {
	for _, c := range implemenetedCMDs {
		if c == cmd {
			return true
		}
	}
	return false
}

type Command struct {
	Length      int
	Cmd         RedisCmd
	Args        []string
	IP          string
	Implemented bool
}

func (c *Command) ToString() string {
	str := string(c.Cmd)
	if len(c.Args) > 0 {
		str += " " + strings.Join(c.Args, " ")
	}
	return str
}

func getCmd(conn *net.TCPConn) (*Command, error) {
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

func makeResStr(cmd *Command) string {
	switch cmd.Cmd {
	case CmdCOMMAND:
		return redisCOMMAND()
	case CmdPING:
		return redisPING()
	case CmdKEYS:
		return redisKEYS()
	case CmdSET:
		return redisSET(cmd.Args)
	case CmdGET:
		return redisGET(cmd.Args[0])
	case CmdDEL:
		return redisDEL(cmd.Args[0])
	case CmdINFO:
		return redisINFO()
	case CmdCONFIG:
		return redisCONFIG()
	case CmdSAVE:
		return redisSAVE()
	}
	return fmt.Sprintf("-ERR unknown command `%s`, with args beginning with: %s", cmd.Cmd, strings.Join(cmd.Args, " "))
}
