package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDbConnection() {
	db, err := gorm.Open(sqlite.Open("myflix.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to db")
	}
	err = db.AutoMigrate(&Video{})
	if err != nil {
		return
	}
	DB = db
}
