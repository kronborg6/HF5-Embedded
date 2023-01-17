package models

import "time"

type Alarm struct {
	Id int `json:"id" gorm:"primaryKey"`
	// LocalName string    `json:"local_id"`
	LocalId     int       `json:"local_id"`
	Local       Startup   `gorm:"foreignKey:LocalId"`
	TimeStamp   time.Time `json:"time_stamp"`
	AlarmTypeId int       `json:"alarm_type_id"`
	AlarmType   AlarmType `gorm:"foreignKey:AlarmTypeId"`
	TypeId      int       `json:"type_id"`
	Type        Types     `gorm:"foreignKey:TypeId"`
	Value       int       `json:"value"`
}

type AlarmType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
