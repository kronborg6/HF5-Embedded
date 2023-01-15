package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/repos"
	"gorm.io/gorm"
)

type AlarmController struct {
	repo *repos.AlarmRepo
}

func (controller *AlarmController) GetAll(c *fiber.Ctx) error {
	alarm, err := controller.repo.FindAll()

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Faild to fetch Users",
			"error":   err,
		})
	}
	return c.JSON(alarm)
}

func (controller *AlarmController) GetById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	alarm, err := controller.repo.FindByID(id)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "can't find alarm by ID",
			"error":   err,
		})
	}

	return c.JSON(&alarm)
}

func NewAlarmController(repo *repos.AlarmRepo) *AlarmController {
	return &AlarmController{repo}
}

func RegisterAlarmController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewAlarmRepo(db)
	controller := NewAlarmController(repo)

	AlarmRouter := router.Group("/alarm")

	AlarmRouter.Get("/", controller.GetAll)
	AlarmRouter.Get("/:id", controller.GetById)
}
