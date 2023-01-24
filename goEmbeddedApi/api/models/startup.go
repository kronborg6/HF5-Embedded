package models

// here we have a the Starup model that the orm use to create a startup table
type Startup struct {
	Id            int     `json:"id" gorm:"primaryKey"`
	Local         string  `json:"local" gorm:"unique"`
	StartTime     int     `json:"start_time"`
	EndTime       int     `json:"end_time"`
	MinTemp       float64 `json:"min_temp"`
	MaxTemp       float64 `json:"max_temp"`
	MinHum        float64 `json:"min_hum"`
	MaxHum        float64 `json:"max_hum"`
	MaxNoiseLevel float64 `json:"max_noise_level"`
}
