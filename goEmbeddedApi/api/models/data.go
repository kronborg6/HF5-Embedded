package models

import "time"

type Data struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	LocalId   int       `json:"local_id"`
	Local     Startup   `gorm:"foreignKey:LocalId"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Hum       float64   `json:"hum"`
	Temp      float64   `json:"temp"`
	// Local     string    `json:"local"`
	// TimeStamp time.Time `json:"time_stamp"`
	// TypeId int     `json:"type_id"`
	// Type   Types   `gorm:"foreignKey:TypeId"`
	// Value  float64 `json:"value"`
}
