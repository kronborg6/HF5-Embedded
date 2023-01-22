package repo

import "gorm.io/gorm"

type DataRepo struct {
	db *gorm.DB
}
