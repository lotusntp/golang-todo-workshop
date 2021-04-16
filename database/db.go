package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
 ID uint `json:"id" gorm:"primary_key"`
 Username string `json:"username"`
 Title string `json:"title"`
 Messgae string `json:"message"`
}

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(sqlite.Open("test.db"),&gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Todo{})

	DB = database
}