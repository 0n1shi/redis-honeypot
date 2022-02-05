package main

const (
	redisNewLine    = "\r\n"
	redisNil        = "$-1" + redisNewLine
	redisMsgOK      = "+OK" + redisNewLine
	redisMsgSuccess = ":1" + redisNewLine
	redisMsgFailure = ":0" + redisNewLine
)

var redisDataMap = map[string]string{}

func redisCOMMAND() string {
	return redisMsgOK
}

func redisPING() string {
	return "+PONG" + redisNewLine
}

func redisKEYS() string { // TODO: return keys
	keys := []string{}
	for k := range redisDataMap {
		keys = append(keys, k)
	}
	return toRedisStrArray(keys)
}

func redisSET(key_value []string) string {
	redisDataMap[key_value[0]] = key_value[1]
	return redisMsgOK
}

func redisGET(key string) string {
	if val, ok := redisDataMap[key]; ok {
		return toRedisStr(val)
	}
	return redisNil
}

func redisDEL(key string) string {
	if _, ok := redisDataMap[key]; !ok {
		return redisMsgFailure
	}
	delete(redisDataMap, key)
	return redisMsgSuccess
}
