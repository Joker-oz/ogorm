package mysql

import (
	"github.com/Joker-oz/ogorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}
func TestNewEngine(t *testing.T) {
	cfg := ogorm.DBConfig{
		DBType: "mysql",
		UserName: "root",
		Password: "DDI@zdns2022",
		Host: "10.2.19.146:3306",
		DBName: "ozweb",
		TablePrefix: "blog_",
		Charset: "utf8",
		ParseTime: true,
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	}
	m := Init()
	_,_ = m.NewEngine(cfg)
	var tag  *Tag
	m.DB.Table("blog_tag").Find(&tag)
	t.Logf("%#v\n",tag)
	m.OpenLog(true)
	m.SetLogger(log.New(os.Stdout, "\r\n", log.LstdFlags),logger.Config{
		SlowThreshold:              time.Second,   // Slow SQL threshold
		IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
	})
	m.DB.Table("blog_tag").Find(&tag)
	t.Logf("%#v\n",tag)
}
