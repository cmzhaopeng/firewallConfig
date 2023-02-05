package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Address struct {
	id           int64
	StartAddress string
	EndAddress   string
	Protocol     string
	Status       int
}

func ConnectDb() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	var address Address
	db.Table("Address").Take(&address)
	fmt.Print(address)

}
