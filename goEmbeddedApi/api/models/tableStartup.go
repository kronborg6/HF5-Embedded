package models

import (
	"time"

	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Types{},
		&Startup{},
		&Data{},
		&Alarm{},
	)
	db.AutoMigrate(
		&Types{},
		&Startup{},
		&Data{},
		&Alarm{},
	)

	startup := []Startup{
		{
			Local:         "Drivehus 1",
			StartTime:     time.Now(),
			EndTime:       time.Now(),
			MinTemp:       5,
			MaxTemp:       15,
			MinHumidity:   10,
			MaxHumidity:   20,
			MaxNoiseLevel: 75,
		},
		{
			Local:         "Drivehus 2",
			StartTime:     time.Now(),
			EndTime:       time.Now(),
			MinTemp:       5,
			MaxTemp:       15,
			MinHumidity:   10,
			MaxHumidity:   20,
			MaxNoiseLevel: 75,
		},
	}
	types := []Types{
		{
			Type: "Sound",
		},
		{
			Type: "Temp",
		},
		{
			Type: "Humidity",
		},
	}
	alarm := []Alarm{
		{
			LocalId:   2,
			TimeStamp: time.Now(),
			TypeId:    1,
			Value:     100,
		},
	}

	db.Create(&startup)
	db.Create(&types)
	db.Create(&alarm)
}
