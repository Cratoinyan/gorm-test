package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	connectionString := ConString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})

	return db, nil
}
