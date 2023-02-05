package main

import (
	"firewallConfig/model"
	"fmt"
)

func main() {
	model.ConnectDb()
	fmt.Println("Hello, playground")
}
