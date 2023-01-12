package repos

import (
	"github.com/kronborg6/HF5-Embedded/goEmbeddedApi/api/models"
	"gorm.io/gorm"
)

type StartupRepo struct {
	db *gorm.DB
}

func (repo *StartupRepo) GetAllStartup() ([]models.Startup, error) {
	var startup []models.Startup

	if err := repo.db.Debug().Find(&startup).Error; err != nil {
		return startup, err
	}

	return startup, nil
}
