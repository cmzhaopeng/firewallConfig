package firewall

import (
	"firewallConfig/model"
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func WritePolicyIpFile(addressList model.AddressList) string {

	var filename string = ""
	var ipGroupName string = ""
	if len(addressList.Addresses) > 0 {
		ipGroupName = addressList.IpGroupName
		filename = filename + string(addressList.Addresses[0].StartIntAddress) + ".txt"
	} else {
		return ""
	}

	//open file name as filename as write the addressList content
	fw, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fw.WriteString("sys")
	fw.WriteString("ip address-set " + ipGroupName + " type object")
	for _, v := range addressList.Addresses {
		fw.WriteString("address range " + v.StartAddress + " " + v.EndAddress)
	}

	fw.WriteString("quit")
	fw.WriteString("quit")
	fw.WriteString("quit")
	return filename
}

func WriteFirewall(filename string) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ip := os.Getenv("FWIP")
	user := os.Getenv("FWUSER")
	password := os.Getenv("FWPASS")
	res := exec.Command("plink.exe", "-ssh", "-l", user, "-pw", password, "-P", "22", "-m", filename, ip, "-sshlog", "permit.log", "-logappend")
	res.Run()
	fmt.Println("Written the address list to the firewall:", ip)

}
