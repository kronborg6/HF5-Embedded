package models

type Types struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Type string `json:"type"`
}
