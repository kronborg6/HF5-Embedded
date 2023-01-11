package models

import "time"

type Alarm struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Local     string    `json:"local"`
	TimeStamp time.Time `json:"time_stamp"`
	WhatType  string    `json:"what_type"`
}
