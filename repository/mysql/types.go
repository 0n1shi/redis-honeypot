package mysql

import (
	"gorm.io/gorm"
)

type Conf struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type RedisCommand struct {
	gorm.Model
	Length      int
	Cmd         string
	Args        string
	IP          string
	Implemented bool
}

func (c *RedisCommand) TableName() string {
	return "redis_commands"
}
