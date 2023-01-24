package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/repos"
	"gorm.io/gorm"
)

// here we say that StartupController is all repos in StartupRepo
type StartupController struct {
	repo *repos.StartupRepo
}

// this get all startup's from the database
func (controller *StartupController) GetAll(c *fiber.Ctx) error {
	// it tell the repo that it want all from the database
	startup, err := controller.repo.GetAllStartup()
	// if it get a error it send a error to the user
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Faild to fetch Users",
			"error":   err,
		})
	}
	// send back all the data
	return c.JSON(startup)
}

// this is a get that get the data whit a id
func (controller *StartupController) GetById(c *fiber.Ctx) error {
	// here it take the id from the url
	id, _ := strconv.Atoi(c.Params("id"))
	// here it send the id to the repo it get back the data if not it get a error
	startup, err := controller.repo.GetByID(id)
	// here it check if it got a error if it godt the error it send a error as a respos
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "can't find a startup whit that id",
			"error":   err,
		})
	}
	// it return the data
	return c.JSON(&startup)
}

// this is the create controller that take in a body
func (controller *StartupController) Create(c *fiber.Ctx) error {
	// here it set what the data must look like
	var startup models.Startup
	var err error
	// here it check if the body look like the model if not it give's a error
	if err = c.BodyParser(&startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// here it send the body data to repo if not it give's a error
	if startup, err = controller.repo.CreateStartup(startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// here it retun the model
	return c.JSON(startup)
}
func Test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Faild to fetch Users",
		"error":   12,
	})
}

// here we have the Update controller that take in a body that is a json body of startup model
func (controller *StartupController) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// here it set what the data must look like
	var startup models.Startup
	var err error
	// here it check if the body look like the model if not it give's a error
	if err = c.BodyParser(&startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// here it send the body data to repo if not it give's a error
	if startup, err = controller.repo.UpdateStartup(id, startup); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// here it retun the model
	return c.JSON(startup)
}

func NewStartupController(repo *repos.StartupRepo) *StartupController {
	return &StartupController{repo}
}

// this create a api group endpoint's the /startup
func RegisterStartupController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewStartupRepo(db)
	controller := NewStartupController(repo)

	StartupRouter := router.Group("/startup")

	StartupRouter.Get("/", controller.GetAll)
	StartupRouter.Get("/:id", controller.GetById)
	StartupRouter.Post("/", controller.Create)
	StartupRouter.Put("/:id", controller.Update)
}
