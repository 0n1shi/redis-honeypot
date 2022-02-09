package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRedisRawClientCmd2Strs(buffer []byte) []string {
	cmdStr := string(buffer)
	strs := strings.Split(cmdStr, "\r\n")
	return strs[:len(strs)-1]
}

func parseRedisClientCmd(strs []string) (*RedisCommand, error) {
	length, err := strconv.Atoi(strs[0][1:])
	if err != nil {
		return nil, err
	}
	cmd := RedisCommand{}
	if length <= 0 {
		return &cmd, nil
	}
	cmd.Length = length
	cmd.Cmd = RedisCmdType(strings.ToUpper(strs[2]))
	for i := 3; i < len(strs); i = i + 2 {
		cmd.Args = append(cmd.Args, strs[i+1])
	}
	return &cmd, nil
}

func toRedisNil() string {
	return "$-1" + redisNewLine
}

func toRedisInt(i int) string {
	return fmt.Sprintf(":%d\r\n", i)
}

func toRedisStr(s string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}

func toRedisStrArray(strs []string) string {
	dataLen := len(strs)
	cmdStr := fmt.Sprintf("*%d\r\n", dataLen)
	for _, s := range strs {
		cmdStr += toRedisStr(s)
	}
	return cmdStr
}

func toRedisIntArray(ints []int) string {
	dataLen := len(ints)
	cmdStr := fmt.Sprintf("*%d\r\n", dataLen)
	for _, i := range ints {
		cmdStr += toRedisInt(i)
	}
	return cmdStr
}
