package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	name  string
	mail  string
	cards []creditCard
}

type creditCard struct {
	gorm.Model
	owner      user
	cardNumber int
	name       string
	expMonth   int
	expYeat    int
	secCode    int
}

func main() {
	connectionString := ConString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&user{})
	db.AutoMigrate(&creditCard{})

}
