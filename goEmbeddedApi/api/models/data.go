package models

import "time"

type Data struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	TimeStamp time.Time `json:"time_stamp"`
}
