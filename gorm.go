package ogorm

import "github.com/Joker-oz/ogorm/mysql"

func NewMysql(config DBConfig) *mysql.Mysql{
	return mysql.Init(config)
}


