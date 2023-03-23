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

func init() {
	mysqlConfig := conf.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Hour)

	var res []User
	db.Table("user").Where("id > 1").Model(&User{}).Scan(&res)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("========", res)
}
