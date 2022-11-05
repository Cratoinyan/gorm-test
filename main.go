package main

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Mail        string
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	UserID     uint
	CardNumber int
	Name       string
	ExpMonth   int
	ExpYeat    int
	SecCode    int
}

func main() {
	db, err := GetDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	_ = db

	fmt.Println("db connected")

}
