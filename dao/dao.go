package dao

import (
	"fmt"
	"goDingTalk/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open(mysql.Open(config.MysqlDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info，打印所有 SQL 语句
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	if Db.Error != nil {
		fmt.Println(Db.Error.Error())
	}

	sqlDB, err := Db.DB()
	if err != nil {
		fmt.Println("mysql connect err is ", err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
