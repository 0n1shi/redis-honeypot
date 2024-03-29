package mysql

import (
	"fmt"
	"strings"

	honeypot "github.com/0n1shi/redis-honeypot"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

var _ honeypot.Repository = (*MySQLRepository)(nil)

func NewMySQLRepository(conf *Conf) (honeypot.Repository, error) {
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
	db.AutoMigrate(&RedisCommand{})

	return &MySQLRepository{db: db}, nil
}

func (r *MySQLRepository) Save(cmd *honeypot.Command) error {
	return r.db.Create(&RedisCommand{
		Length:      cmd.Length,
		Cmd:         string(cmd.Cmd),
		Args:        strings.Join(cmd.Args, " "),
		IPFrom:      cmd.IPFrom,
		IPTo:        cmd.IPTo,
		Implemented: cmd.Implemented,
	}).Error
}
