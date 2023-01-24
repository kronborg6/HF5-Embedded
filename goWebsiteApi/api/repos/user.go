package repos

import (
	"errors"
	"fmt"

	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/middleware"
	"github.com/kronborg6/HF5-Embedded/goWebsiteApi/api/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// username string, pass string
func (repo *UserRepo) FindUser(data models.User) (*[]models.User, error) {
	var user []models.User

	err := repo.db.Find(&user).Error

	if err != nil {
		return nil, err
	}

	// fmt.Println(user[0].Password)
	if user[0].Password != middleware.Encode(data.Password) {
		return nil, nil
	}
	fmt.Println("user")
	fmt.Println(&user)

	return &user, nil
}

func (repo *UserRepo) FindUserByEmail(username string) (models.User, error) {
	var user models.User

	q := repo.db.Where("username = ?", username).Find(&user)

	if q.Error != nil {
		return user, q.Error
	}
	if q.RowsAffected <= 0 {
		return user, errors.New("can't find user")
	}

	return user, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
