package models

// this model is only to show the orm what it need

type Types struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Type string `json:"type"`
}
