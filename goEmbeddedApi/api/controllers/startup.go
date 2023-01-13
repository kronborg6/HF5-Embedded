package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/repos"
	"gorm.io/gorm"
)

type StartupController struct {
	repo *repos.StartupRepo
}

func (controller *StartupController) GetAll(c *fiber.Ctx) error {
	startup, err := controller.repo.GetAllStartup()

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Faild to fetch Users",
			"error":   err,
		})
	}
	return c.JSON(startup)
}

func NewStartupController(repo *repos.StartupRepo) *StartupController {
	return &StartupController{repo}
}

func RegisterStartupController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewStartupRepo(db)
	controller := NewStartupController(repo)

	StartupRouter := router.Group("/startup")

	StartupRouter.Get("/", controller.GetAll)
}
