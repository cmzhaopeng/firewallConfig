package model

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

func (Address) TableName() string {
	return "Address"
}

func (IpGroup) TableName() string {
	return "IpGroup"
}

func (AddressIpGroup) TableName() string {
	return "AddressIpGroup"
}

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

}

func QueryAddress() AddressList {
	var addresses []Address
	result := DB.Limit(10).Where("status = ? AND address_type= ? AND end_int_address-start_int_address < ?", 1, "W", 256).Find(&addresses)
	if result.Error != nil {
		fmt.Println(result.Error)
		return AddressList{}
	}

	var addCount int64 = result.RowsAffected
	if addCount == 0 {
		return AddressList{}
	}

	var ipGroup IpGroup

	result = DB.Where("ip_count< ?", 3000).First(&ipGroup)
	if result.Error != nil {
		fmt.Println(result.Error)
		return AddressList{}
	}

	name := ipGroup.Name
	ipGroup.IpCount = ipGroup.IpCount + int(addCount)
	result = DB.Model(&ipGroup).Where("name = ?", name).Update("ip_count", ipGroup.IpCount)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	var addressList AddressList

	for i := 0; i < len(addresses); i++ {
		addresses[i].Status = 2
		result = DB.Model(&addresses[i]).Where("id = ?", addresses[i].ID).Update("status", addresses[i].Status)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		addressList.Addresses = append(addressList.Addresses, addresses[i])
		addressList.IpGroupName = name
		DB.Create(&AddressIpGroup{AddressId: addresses[i].ID, IpGroupId: ipGroup.ID})
	}

	return addressList
}
