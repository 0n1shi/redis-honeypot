package mysql

import (
	"fmt"
	"strings"

	"github.com/0n1shi/redis-honeypot/pkg/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

var _ redis.Repository = (*MySQLRepository)(nil)

func NewMySQLRepository(conf *Conf) (redis.Repository, error) {
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

func (r *MySQLRepository) Save(cmd *redis.Command) error {
	return r.db.Create(&RedisCommand{
		Length:      cmd.Length,
		Cmd:         string(cmd.Cmd),
		Args:        strings.Join(cmd.Args, " "),
		IP:          cmd.IP,
		Implemented: cmd.Implemented,
	}).Error
}
