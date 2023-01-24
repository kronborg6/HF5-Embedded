package repos

import (
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
	"gorm.io/gorm"
)

type StartupRepo struct {
	db *gorm.DB
}

// this gte all the startup from the database
func (repo *StartupRepo) GetAllStartup() ([]models.Startup, error) {
	// we set the model
	var startup []models.Startup
	// it try's to get all startup from databasen if not it retun a error
	if err := repo.db.Debug().Find(&startup).Error; err != nil {
		return startup, err
	}
	// retun all the data form databasen
	return startup, nil
}

// git startup by id from the database
func (repo *StartupRepo) GetByID(id int) (*[]models.Startup, error) {
	//it set the model it gona use
	var startup []models.Startup
	// tells the data give me were id is and the table match startup model
	err := repo.db.Debug().Where("id = ?", id).Find(&startup).Error
	// if error it return a error
	if err != nil {
		return &startup, err
	}
	// retun the data from the database
	return &startup, nil
}

// crate startup in the database
func (repo *StartupRepo) CreateStartup(startup models.Startup) (models.Startup, error) {
	// it trys to create startup where the model matchs
	err := repo.db.Create(&startup).Error
	// if error it retun a error
	if err != nil {
		return startup, err
	}
	// retun the data
	return startup, nil
}

// Update startup
func (repo *StartupRepo) UpdateStartup(id int, startup models.Startup) (models.Startup, error) {

	// its update the the startup only where the data was change
	err := repo.db.Model(&startup).Where("id = ?", id).Updates(&startup).Error
	// if it get a error it retun a error
	if err != nil {
		return startup, err
	}
	// return the update data
	return startup, nil
}

// retun the repo whit the database

func NewStartupRepo(db *gorm.DB) *StartupRepo {
	return &StartupRepo{db}
}
