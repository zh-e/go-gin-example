package models

import (
	"gorm.io/gorm"
	"log"
)

type UserMoney struct {
	ID     uint64
	userId string
	gold   string
}

func (UserMoney) TableName() string {
	return "app_user_money"
}

func (UserMoney) AfterSave(db *gorm.DB) (err error) {
	log.Println("===========", "UserMoney AfterSave")
	return
}
