package models

import (
	"fmt"
	"go-gin-example/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

// https://gorm.io/zh_CN/docs/models.html
func init() {
	mysqlConfig := conf.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)

	var err error
	// 此处不可使用 := 给db赋值 要不然db只会是局部可用的
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("数据库连接初始化完成")
}
