package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

func CreateDB(){
	db, err := gorm.Open(sqlite.Open("test.db"),&gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})
}