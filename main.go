package main

import (
	"firewallConfig/firewall"
	"firewallConfig/model"
	"fmt"
)

func main() {
	model.ConnectDb()

	// for {
	var addressList model.AddressList = model.QueryAddress()
	for len(addressList.Addresses) > 0 {

		//fmt.Print(addressList)
		filename := firewall.WritePolicyIpFile(addressList)
		fmt.Println(filename)
		/*
			if filename!="" {
				firewall.WriteFirewall(filename)
			}
		*/
		addressList = model.QueryAddress()
	}
	//time.Sleep(timeout * time.Minute)
	//}
}
