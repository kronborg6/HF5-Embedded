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
	// start the databae conniton
	db := db.Init()
	// make's a new api
	app := fiber.New()
	// startup off models
	models.Setup(db)
	// log's the all the req
	app.Use(logger.New())
	// a endpoint test
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this is a endpoint test",
			"test":    "Test",
		})
	})
	// setup of basic auth
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"Admin": "Password",
		},
	}))
	// tells what the endpoint is
	api := app.Group("/")
	controllers.RegisterStartupController(db, api)
	controllers.RegisterAlarmController(db, api)
	controllers.RegisterDataController(db, api)
	// listen to port 8080 for the api
	log.Fatal(app.Listen(":8080"))
}
