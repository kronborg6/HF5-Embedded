package models

import (
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Types{},
		&Startup{},
		&Data{},
		&AlarmType{},
		&Alarm{},
	)
	db.AutoMigrate(
		&Types{},
		&Startup{},
		&Data{},
		&AlarmType{},
		&Alarm{},
	)

	startup := []Startup{
		{
			Local:         "Drivehus 1",
			StartTime:     8,
			EndTime:       18,
			MinTemp:       5,
			MaxTemp:       100,
			MinHum:        10,
			MaxHum:        20,
			MaxNoiseLevel: 75,
		},
		{
			Local:         "Drivehus 2",
			StartTime:     8,
			EndTime:       16,
			MinTemp:       5,
			MaxTemp:       15,
			MinHum:        10,
			MaxHum:        20,
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
	alarmType := []AlarmType{
		{
			Name: "Warning",
		},
		{
			Name: "Alarm",
		},
	}
	alarm := []Alarm{
		{
			LocalId: 2,
			// TimeStamp:   time.Now(),
			AlarmTypeId: 1,
			TypeId:      1,
			Value:       100,
		},
	}
	data := []Data{
		{
			LocalId: 1,
			// TypeId:  1,
			// Value:   25,
			Temp: 21.2,
			Hum:  15.2,
		},
	}

	db.Create(&startup)
	db.Create(&types)
	db.Create(&data)
	db.Create(&alarmType)
	db.Create(&alarm)
}
