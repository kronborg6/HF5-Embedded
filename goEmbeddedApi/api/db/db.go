package db

import (
	"log"
	// "github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "dfg"
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
