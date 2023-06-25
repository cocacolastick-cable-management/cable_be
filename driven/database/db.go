package database

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) (db *gorm.DB) {

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	db = db.Debug()

	db.Raw("CREATE EXTENSION IF NOT EXISTS 'uuid-ossp'", nil)

	err = db.AutoMigrate(&entities.User{}, &entities.Contract{}, &entities.Request{}, &entities.RequestHistory{}, &entities.Notification{})
	if err != nil {
		panic(err)
	}

	return db
}
