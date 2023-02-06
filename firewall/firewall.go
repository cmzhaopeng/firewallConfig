package firewall

import (
	"firewallConfig/model"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/joho/godotenv"
)

func WritePolicyIpFile(addressList model.AddressList) string {

	var filename string = ""
	var ipGroupName string = ""
	if len(addressList.Addresses) > 0 {
		ipGroupName = addressList.IpGroupName
		filename = "cmd" + strconv.FormatInt(addressList.Addresses[0].StartIntAddress, 10) + ".txt"
	} else {
		return ""
	}

	//open file name as filename as write the addressList content
	fw, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fw.WriteString("sys\n")
	fw.WriteString("ip address-set " + ipGroupName + " type object\n")
	for _, v := range addressList.Addresses {
		if v.StartAddress == v.EndAddress {
			fw.WriteString("address " + v.StartAddress + " mask 32\n")
		} else {
			fw.WriteString("address range " + v.StartAddress + " " + v.EndAddress + "\n")
		}

	}

	fw.WriteString("quit\n")
	fw.WriteString("quit\n")
	fw.WriteString("quit\n")
	fw.Close()
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
