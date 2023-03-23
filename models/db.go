package models

import (
	"fmt"
	"go-gin-example/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitDB() {
	mysqlConfig := conf.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)

	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

}
