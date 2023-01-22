package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/db"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
)

func main() {
	db := db.Init()
	app := fiber.New()
	models.Setup(db)
	/*
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		})) */
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this is a endpoint test",
			"test":    "Test",
		})
	})

	log.Fatal(app.Listen(":8000"))

}
