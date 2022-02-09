package main

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLRedisCommand struct {
	gorm.Model
	Length      int
	Cmd         string
	Args        string
	Addr        string
	Implemented bool
}

func (c *MySQLRedisCommand) TableName() string {
	return "redis_commands"
}

func initMySQL(conf *MySQLConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&MySQLRedisCommand{})

	return db, nil
}

func toMySQLRecord(cmd *RedisCommand) *MySQLRedisCommand {
	return &MySQLRedisCommand{
		Length:      cmd.Length,
		Cmd:         string(cmd.Cmd),
		Args:        strings.Join(cmd.Args, " "),
		Addr:        cmd.Addr,
		Implemented: cmd.Implemented,
	}
}
