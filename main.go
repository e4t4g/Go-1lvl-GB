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
	fmt.Println("DB_URL: ", config.Db_url)
	fmt.Println("some_app_key: ", config.Some_app_key)
}
