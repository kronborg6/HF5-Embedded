package main

import (
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/db"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
)

func main() {
	db := db.Init()

	models.Setup(db)
}
