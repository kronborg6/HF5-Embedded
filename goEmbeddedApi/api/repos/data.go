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

	if err := repo.db.Debug().Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewDataRepo(db *gorm.DB) *DataRepo {
	return &DataRepo{db}
}
