package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goDingTalk/config"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.MysqlDB)
	if err != nil {
		fmt.Println(err.Error())
	}
	if Db.Error != nil {
		fmt.Println(Db.Error.Error())
	}

	Db.LogMode(true) // 启用SQL日志记录

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}
