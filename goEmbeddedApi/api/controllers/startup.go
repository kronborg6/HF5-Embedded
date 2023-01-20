package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
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

func (controller *StartupController) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	startup, err := controller.repo.GetByID(id)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "can't find a startup whit that id",
			"error":   err,
		})
	}

	return c.JSON(&startup)
}

func (controller *StartupController) Create(c *fiber.Ctx) error {
	var startup models.Startup
	var err error

	if err = c.BodyParser(&startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if startup, err = controller.repo.CreateStartup(startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(startup)
}

func (controller *StartupController) Update(c *fiber.Ctx) error {
	var startup models.Startup
	var err error
	// id, _ := strconv.Atoi(c.Params("id"))

	if err = c.BodyParser(&startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if startup, err = controller.repo.UpdateStartup(startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
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
	StartupRouter.Get("/:id", controller.GetById)
	StartupRouter.Post("/", controller.Create)
	StartupRouter.Put("/", controller.Update)
}
