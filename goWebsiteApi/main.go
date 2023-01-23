package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/controllers"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/db"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/middleware"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
)

var (
	// Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	privateKey = "dfg"
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

		/* 		return c.JSON(fiber.Map{
			"message": "this is a endpoint test",
			"test":    "Test",
		}) */
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	app.Get("/no", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "fuck no",
		})
	})

	fmt.Println(middleware.Encode("Kronborg"))
	fmt.Println(middleware.Dcode("S3JvbmJvcmc="))

	api := app.Group("/")

	controllers.RegisterDataController(db, api)

	log.Fatal(app.Listen(":8000"))

}
