package mysql

import (
	"fmt"
	"github.com/Joker-oz/ogorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Model struct {
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

type Mysql struct {
	logger    logger.Interface
	DB        *gorm.DB
	isOpenLog bool
}

func Init(config ogorm.DBConfig) *Mysql {
	m := &Mysql{}
	_,err := m.NewEngine(config)
	if err != nil {
		log.Println("ogorm: init mysql db fail!")
	}
	return m
}

func (m *Mysql) NewEngine(cfg ogorm.DBConfig) (*gorm.DB, error) {
	conn := fmt.Sprintf(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.DBName,
		cfg.Charset,
		cfg.ParseTime,
	))
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,
		SkipInitializeWithVersion: cfg.SkipInitializeWithVersion,
	}))
	m.DB = db
	if cfg.Log {
		m.isOpenLog = true
		m.SetDefaultLogger()
	}
	return db, err
}

func (m *Mysql) OpenLog(open bool) {
	m.isOpenLog = open
	if open {
		if m.logger == nil {
			log.Println("ogrom: current mysql db logger is nil, please set logger")
		}
		m.DB.Logger = m.logger
	} else {
		m.DB.Logger = nil
	}
}

func (m *Mysql) SetLogger(writer logger.Writer, config logger.Config) {
	config = m.mergeDefaultLogCfg(config)
	m.logger = logger.New(writer, config)
	if m.isOpenLog {
		m.DB.Logger = m.logger
	}
}

func (m *Mysql) SetDefaultLogger() {
	m.SetLogger(log.New(os.Stdout, "\r\n", log.LstdFlags),ogorm.DefaultDBLogConfig)
}

func (m *Mysql) mergeDefaultLogCfg(config logger.Config) logger.Config {
	if config.SlowThreshold == 0 {
		config.SlowThreshold = ogorm.DefaultDBLogConfig.SlowThreshold
	}
	if config.LogLevel == 0 {
		config.LogLevel = ogorm.DefaultDBLogConfig.LogLevel
	}
	if config.Colorful == false {
		config.Colorful = ogorm.DefaultDBLogConfig.Colorful
	}
	if config.IgnoreRecordNotFoundError == false {
		config.IgnoreRecordNotFoundError = ogorm.DefaultDBLogConfig.IgnoreRecordNotFoundError
	}
	return config
}
