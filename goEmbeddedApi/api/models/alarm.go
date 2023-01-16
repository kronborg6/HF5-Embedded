package models

import "time"

type Alarm struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	LocalName string    `json:"local_id"`
	Local     Startup   `gorm:"foreignKey:LocalName"`
	TimeStamp time.Time `json:"time_stamp"`
	TypeId    int       `json:"type_id"`
	Type      Types     `gorm:"foreignKey:TypeId"`
	Value     int       `json:"value"`
}
