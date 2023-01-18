package db

import (
	"log"
	// "github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dsn := "kronborg:password@tcp(localhost:3306)/lol?charset=utf8mb4&parseTime=True&loc=Local" // Home Database
	dsn := "root:Password@tcp(localhost:3306)/lol?charset=utf8mb4&parseTime=True&loc=Local" // connect to wsl mysql server
	// dsn := "mysql://example_user:password@127.0.0.1:3306/exampleDB?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:@tcp(127.0.0.1:3306)/lol?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Setyup(db *gorm.DB) {
	db.AutoMigrate()
}
