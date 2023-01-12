package models

import "time"

type Data struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Local     string    `json:"local"`
	TimeStamp time.Time `json:"time_stamp"`
	TypeId    int       `json:"type_id"`
	Type      Types     `gorm:"foreignKey:TypeId"`
	Value     int       `json:"value"`
}
