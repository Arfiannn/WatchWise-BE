package main

import (
	"watchwise_be/config"
	"watchwise_be/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}