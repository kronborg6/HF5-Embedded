package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/db"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
)

func main() {
	db := db.Init()
	app := fiber.New()

	models.Setup(db)

	// fmt.Println(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("LoL")
	})

	app.Listen(":8080")
}
