package models

import "time"

type Alarm struct {
	Id int `json:"id" gorm:"primaryKey"`
	// LocalName string    `json:"local_id"`
	LocalId int     `json:"local_id"`
	Local   Startup `gorm:"foreignKey:LocalId"`
	// TimeStamp   time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
	AlarmTypeId int       `json:"alarm_type_id"`
	AlarmType   AlarmType `gorm:"foreignKey:AlarmTypeId"`
	TypeId      int       `json:"type_id"`
	Type        Types     `gorm:"foreignKey:TypeId"`
	Value       int       `json:"value"`
	CreatedAt   time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
}

type AlarmType struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
