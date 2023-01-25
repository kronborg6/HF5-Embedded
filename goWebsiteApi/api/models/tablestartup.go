package models

import "gorm.io/gorm"

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&User{},
	)
	db.AutoMigrate(
		&User{},
	)

	user := []User{
		{
			Username: "Kronborg",
			Password: "S3JvbmJvcmc=",
		},
		{
			Username: "Bob",
			Password: "Qm9i",
		},
	}

	db.Create(&user)
}
