package redishoneypot

import (
	"fmt"
	"net"
	"strings"
)

const (
	CmdCOMMAND  = "COMMAND"
	CmdPING     = "PING"
	CmdKEYS     = "KEYS"
	CmdSET      = "SET"
	CmdGET      = "GET"
	CmdDEL      = "DEL"
	CmdINFO     = "INFO"
	CmdCONFIG   = "CONFIG"
	CmdSAVE     = "SAVE"
	CmdQUIT     = "QUIT"
	CmdFLUSHALL = "FLUSHALL"
	CmdClient   = "CLIENT"
	CmdSLAVEOF  = "SLAVEOF"
	CmdAUTH     = "AUTH"
	CmdHELP     = "HELP"
)

var implemenetedCmds = map[string](func(args []string) string){
	CmdCOMMAND:  redisCOMMAND,
	CmdPING:     redisPING,
	CmdKEYS:     redisKEYS,
	CmdSET:      redisSET,
	CmdGET:      redisGET,
	CmdDEL:      redisDEL,
	CmdINFO:     redisINFO,
	CmdCONFIG:   redisCONFIG,
	CmdSAVE:     redisSAVE,
	CmdQUIT:     redisQUIT,
	CmdFLUSHALL: redisFLUSHALL,
	CmdClient:   redisCLIENT,
	CmdSLAVEOF:  redisSLAVEOF,
	CmdAUTH:     redisAUTH,
	CmdHELP:     redisHELP,
}

func IsImplemented(cmd string) bool {
	_, ok := implemenetedCmds[cmd]
	return ok
}

type Command struct {
	Length      int
	Cmd         string
	Args        []string
	IPFrom      string
	IPTo        string
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
	strs := parseRawCmdToStrs(buffer)
	cmd, err := parseStrsToClientCmd(strs)
	if err != nil {
		return nil, err
	}
	cmd.IPFrom = strings.Split(conn.RemoteAddr().String(), ":")[0]
	cmd.IPTo = strings.Split(conn.LocalAddr().String(), ":")[0]
	cmd.Implemented = IsImplemented(cmd.Cmd)
	return cmd, nil
}

func makeResStr(cmd *Command) string {
	if function, ok := implemenetedCmds[cmd.Cmd]; ok {
		return function(cmd.Args)
	}
	return fmt.Sprintf("-ERR unknown command `%s`, with args beginning with: %s%s", cmd.Cmd, strings.Join(cmd.Args, " "), ResNewLine)
}
