package repos

import (
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
	"gorm.io/gorm"
)

type AlarmRepo struct {
	db *gorm.DB
}

func (repo *AlarmRepo) FindAll() ([]models.Alarm, error) {
	var alarm []models.Alarm

	if err := repo.db.Debug().Find(&alarm).Error; err != nil {
		return alarm, err
	}

	return alarm, nil
}

func NewAlarmRepo(db *gorm.DB) *AlarmRepo {
	return &AlarmRepo{db}
}
