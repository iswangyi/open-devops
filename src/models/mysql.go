package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"open-devops/src/modules/server/config"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var DB = map[string]*xorm.Engine{}

// MySQLInit 初始化数据库连接
func MySQLInit(sConf []*config.MySQLConf) error {
	for _, conf := range sConf {
		db, err := xorm.NewEngine("mysql", conf.Addr)
		if err != nil {
			fmt.Printf("[%v:] 链接失败,[err:%v]\n", conf.Addr, err)
			continue
		}
		db.SetMaxIdleConns(conf.Idle)
		db.SetMaxOpenConns(conf.Max)
		db.SetConnMaxLifetime(time.Hour)
		db.ShowSQL(conf.Debug)
		db.Logger().SetLevel(xlog.LOG_INFO)
		DB[conf.Name] = db
	}
	return nil
}
