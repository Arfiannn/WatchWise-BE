package main

import (
	"fmt"
	"watchwise_be/config"
	"watchwise_be/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	fmt.Println("🚀 WatchWise API berjalan di http://localhost:8080 🚀")
	r.Run(":8080")
}
