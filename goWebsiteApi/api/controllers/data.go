package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/repos"
	"gorm.io/gorm"
)

type DataController struct {
	repo *repos.DataRepo
}

func (controller *DataController) GetAll(c *fiber.Ctx) error {
	data, err := controller.repo.FindAllDta()

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "faild to fetch data",
			"error":   err,
		})
	}
	return c.JSON(data)
}

func NewDAtaController(repo *repos.DataRepo) *DataController {
	return &DataController{repo}
}
func RegisterDataController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewDataRepo(db)
	controller := NewDAtaController(repo)

	DataRouter := router.Group("/data")

	DataRouter.Get("/", controller.GetAll)
}
