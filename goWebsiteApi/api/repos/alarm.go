package repos

import "gorm.io/gorm"

type AlarmRepo struct {
	db *gorm.DB
}

func NewAlarmRepo(db *gorm.DB) *AlarmRepo {
	return &AlarmRepo{db}
}
