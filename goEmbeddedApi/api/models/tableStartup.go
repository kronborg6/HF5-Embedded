package models

import "gorm.io/gorm"

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Types{},
		&Alarm{},
		&Startup{},
		&Data{},
	)
	db.AutoMigrate(
		&Types{},
		&Alarm{},
		&Startup{},
		&Data{},
	)
}
