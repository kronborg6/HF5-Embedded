package models

// this model is only to show the orm what it need

import "time"

type Data struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	LocalId   int       `json:"local_id"`
	Local     Startup   `gorm:"foreignKey:LocalId"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	Hum       float64   `json:"hum"`
	Temp      float64   `json:"temp"`
}
