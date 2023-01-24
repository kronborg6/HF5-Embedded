package repos

import (
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
	"gorm.io/gorm"
)

type DataRepo struct {
	db *gorm.DB
}

func (repo *DataRepo) FindAllData() ([]models.Data, error) {
	var data []models.Data

	if err := repo.db.Debug().Preload("Local").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewDataRepo(db *gorm.DB) *DataRepo {
	return &DataRepo{db}
}
