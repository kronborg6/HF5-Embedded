package repos

import (
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) FindUser() (*[]models.User, error) {

}
