package main

import (
	"firewallConfig/firewall"
	"firewallConfig/model"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		fmt.Println(err)
	}

	model.ConnectDb()
	for {

		var addressList model.AddressList = model.QueryAddress()
		for len(addressList.Addresses) > 0 {

			//fmt.Print(addressList)
			filename := firewall.WritePolicyIpFile(addressList)
			fmt.Println(filename)

			if filename != "" {
				firewall.WriteFirewall(filename)
			}

			addressList = model.QueryAddress()
		}
		time.Sleep(time.Minute * time.Duration(timeout))
	}

}
