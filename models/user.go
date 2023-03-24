package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID         int64
	Name       string
	Sex        string
	CreateTime time.Time
	UpdateTime time.Time
}

func Create(name string, sex string) {
	user := User{Name: name, Sex: sex}
	db.Model(&User{}).Create(&user)
}

func Find(id int64) User {
	var user User

	db.Model(&User{}).Find(&user, id)

	return user
}

// TableName gorm 默认model转表 采用的是 model名字+复数 自定义表名需要实现此接口
func (u *User) TableName() string {
	return "app_user"
}

func (u *User) AfterSave(db *gorm.DB) (err error) {
	log.Println("===========", "user AfterSave")
	return
}
