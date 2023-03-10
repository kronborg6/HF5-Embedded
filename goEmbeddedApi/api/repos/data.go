package repos

import (
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
	"gorm.io/gorm"
)

type DataRepo struct {
	db *gorm.DB
}

func (repo *DataRepo) FindAll() ([]models.Data, error) {
	var data []models.Data

	if err := repo.db.Debug().Preload("Local").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (repo *DataRepo) FindByID(id int) (*[]models.Data, error) {
	var data []models.Data

	err := repo.db.Debug().Preload("Type").Where("id = ?", id).Find(&data).Error

	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (repo *DataRepo) CreateData(data models.Data) (models.Data, error) {
	err := repo.db.Create(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}
func (repo *DataRepo) DeleteData(data models.Data) error {

	err := repo.db.Debug().Delete(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func NewDataRepo(db *gorm.DB) *DataRepo {
	return &DataRepo{db}
}
