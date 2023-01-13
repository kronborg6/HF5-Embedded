package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/controllers"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/db"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
)

func main() {
	db := db.Init()
	app := fiber.New()

	models.Setup(db)

	// fmt.Println(db)

	app.Get("/lol", func(c *fiber.Ctx) error {
		return c.SendString("LoL")
	})

	api := app.Group("/")
	controllers.RegisterStartupController(db, api)
	controllers.RegisterAlarmController(db, api)
	controllers.RegisterDataController(db, api)

	app.Listen(":8080")
}
