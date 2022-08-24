package main

import (
	"gin-crud/handlers"
	"gin-crud/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":8080")
}
