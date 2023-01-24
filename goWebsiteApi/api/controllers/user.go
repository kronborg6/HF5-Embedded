package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/middleware"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/repos"
	"gorm.io/gorm"
)

type UserController struct {
	repo *repos.UserRepo
}

func (controller *UserController) test(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	data, err := controller.repo.FindUserByEmail(user.Username)
	// fmt.Println("Error: " + err.Error())
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "can't find data by ID",
			"error":   err,
		})
	}
	if data.Password != middleware.Encode(user.Password) {
		return c.JSON(fiber.Map{
			"message": "password or username do not match",
		})
	}

	claims := jwt.MapClaims{
		"name": user.Username,
		"exp":  time.Now().Add(time.Hour * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}

func (controller *UserController) login(c *fiber.Ctx) error {
	username := c.FormValue("pass")
	var user models.User

	fmt.Println("her is the username: " + username)
	data, err := controller.repo.FindUser(user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "can't find data by ID",
			"error":   err,
		})
	}
	fmt.Println(data)

	// return c.JSON(fiber.Map{
	// 	"message": "Hej med dig jeg heder kaj",
	// })
	return c.JSON(&data)
}

func NewUserController(repo *repos.UserRepo) *UserController {
	return &UserController{repo}
}

func RegisterUserController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewUserController(repo)

	UserRouter := router.Group("/user")

	UserRouter.Post("/login", controller.login)
	UserRouter.Post("/", controller.test)
}
