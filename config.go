package ogorm

import (
	"gorm.io/gorm/logger"
	"time"
)

type DBConfig struct {
	DBType                    string
	UserName                  string
	Password                  string
	Host                      string
	DBName                    string
	TablePrefix               string
	Charset                   string
	ParseTime                 bool
	MaxIdleConns              int
	MaxOpenConns              int
	Log                       bool
	SkipInitializeWithVersion bool // 跳过数据库初始配置：禁用索引、表名等修改
}


var DefaultDBLogConfig = logger.Config{
	SlowThreshold:             time.Second,
	LogLevel:                  logger.Info,
	IgnoreRecordNotFoundError: true,
	Colorful:                  true,
}
