package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// here we connet to the database useing a conntion string
func Init() *gorm.DB {
	// dsn := "kronborg:password@tcp(localhost:3306)/lol?charset=utf8mb4&parseTime=True&loc=Local" // Home Database
	dsn := "kronborg:password@tcp(localhost:3306)/lol?charset=utf8mb4&parseTime=True&loc=Local" // Home Database

	// dsn := "root:Password@tcp(localhost:3306)/lol?charset=utf8mb4&parseTime=True&loc=Local" // connect to wsl mysql server
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Setyup(db *gorm.DB) {
	db.AutoMigrate()
}
