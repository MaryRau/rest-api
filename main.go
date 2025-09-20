package main

import (
	"fmt"
	"restapi/database"
	"restapi/http"
)

func init() {
	database.Connect()
}

func main() {
	if err := http.StartServer(); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
