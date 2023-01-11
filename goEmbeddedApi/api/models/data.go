package models

import "time"

type Data struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Local     string    `json:"local"`
	TimeStamp time.Time `json:"time_stamp"`
	Temp      int       `json:"temp"`
	Humidity  int       `json:"humidity"`
}
