package main

import (
	"Go-1lvl-GB/config"
	"fmt"
)

func main() {
	config, err := config.CreateNew()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("port: ", config.Port)
	fmt.Println("DbUrl: ", config.DbUrl)
	fmt.Println("someAppKey: ", config.SomeAppKey)
}
