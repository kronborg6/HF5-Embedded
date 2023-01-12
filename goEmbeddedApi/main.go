package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/db"
)

func main() {
	db := db.Init()

	fmt.Println(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("LoL")
	})

	app.Listen(":8080")
}
