package ogorm

func NewMysql(config DBConfig) *Mysql {
	return Init(config)
}


