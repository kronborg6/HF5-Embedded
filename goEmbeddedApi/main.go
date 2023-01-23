package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/controllers"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/db"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
)

func main() {
	db := db.Init()
	app := fiber.New()

	models.Setup(db)

	// fmt.Println(db)

	app.Use(logger.New())

	app.Get("/lol", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "can't find data by ID",
			"Nej":     "Hej",
		})
	})
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"Lort": "Password",
		},
	}))

	api := app.Group("/")
	controllers.RegisterStartupController(db, api)
	controllers.RegisterAlarmController(db, api)
	controllers.RegisterDataController(db, api)
	log.Fatal(app.Listen(":8080"))
	// app.Listen(":8080")
}
