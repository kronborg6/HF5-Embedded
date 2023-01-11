package models

import "time"

type Startup struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	Local         string    `json:"local"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	MinTemp       int       `json:"min_temp"`
	MaxTemp       int       `json:"max_temp"`
	MinHumidity   int       `json:"min_humidity"`
	MaxHumidity   int       `json:"max_humidty"`
	MaxNoiseLevel int       `json:"max_noise_level"`
}
