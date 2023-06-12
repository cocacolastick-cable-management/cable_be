package database

import (
	"github.com/cable_management/cable_be/app/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init(dsn string) {

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	DB = DB.Debug()

	err = DB.AutoMigrate(&entities.User{})
	if err != nil {
		panic(err)
		return
	}
}
