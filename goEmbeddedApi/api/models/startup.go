package models

// here we have a the Starup model that the orm use to create a startup table
type Startup struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Local         string `json:"local" gorm:"unique"`
	StartTime     int    `json:"start_time"`
	EndTime       int    `json:"end_time"`
	MinTemp       int    `json:"min_temp"`
	MaxTemp       int    `json:"max_temp"`
	MinHum        int    `json:"min_hum"`
	MaxHum        int    `json:"max_hum"`
	MaxNoiseLevel int    `json:"max_noise_level"`
}
