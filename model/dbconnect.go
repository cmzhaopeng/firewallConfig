package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Address struct {
	ID              int64
	StartAddress    string
	EndAddress      string
	Protocol        string
	StartIntAddress int64
	EndIntAddress   int64
	Status          int
}

type IpGroup struct {
	ID      int64
	Name    string
	IpCount int
}

type AddressIpGroup struct {
	ID        int64
	AddressId int64
	IpGroupId int64
}

type AddressList struct {
	Addresses   []Address
	IpGroupName string
}

var DB *gorm.DB

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
	DB = db
	var address Address
	var ipGroup IpGroup
	var addressIpGroup AddressIpGroup
	db.Table("Address").Take(&address)
	db.Table("IpGroup").Take(&ipGroup)
	db.Table("AddressIpGroup").Take(&addressIpGroup)

	fmt.Print(address)
	fmt.Print(ipGroup)
	fmt.Print(addressIpGroup)
}

func QueryAddress() AddressList {
	var address Address
	DB.Table("Address").Take(&address)
	var addressList AddressList
	addressList.Addresses = append(addressList.Addresses, address)
	addressList.IpGroupName = "cg-6"
	return addressList
}
