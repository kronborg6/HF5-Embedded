package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/repos"
	"gorm.io/gorm"
)

type DataController struct {
	repo *repos.DataRepo
}

func (controller *DataController) GetAll(c *fiber.Ctx) error {
	data, err := controller.repo.FindAll()

	if err != nil {
		return c.JSON((fiber.Map{
			"message": "faild to fetch data",
			"error":   err,
		}))
	}
	return c.JSON(data)
}

func (controller *DataController) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	data, err := controller.repo.FindByID(id)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "can't find data by ID",
			"error":   err,
		})
	}
	return c.JSON(&data)
}

func (controller *DataController) Create(c *fiber.Ctx) error {
	var data models.Data
	var err error

	if err = c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if data, err = controller.repo.CreateData(data); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

func NewDataController(repo *repos.DataRepo) *DataController {
	return &DataController{repo}
}

func RegisterDataController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewDataRepo(db)
	controller := NewDataController(repo)

	DataRouter := router.Group("/data")

	DataRouter.Get("/", controller.GetAll)
	DataRouter.Get("/:id", controller.GetById)
}
