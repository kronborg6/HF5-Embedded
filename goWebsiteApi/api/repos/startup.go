package repos

import "gorm.io/gorm"

type StartupRepo struct {
	db *gorm.DB
}

func NewStartupRepo(db *gorm.DB) *StartupRepo {
	return &StartupRepo{db}
}
