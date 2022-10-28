package ogorm

import (
	"github.com/Joker-oz/ogorm/config"
	"github.com/Joker-oz/ogorm/mysql"
)

func NewMysql(config config.DBConfig) *mysql.Mysql {
	return mysql.Init(config)
}


