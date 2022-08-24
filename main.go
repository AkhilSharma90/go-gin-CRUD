package main

import (
	"go-gin-CRUD/handlers"
	"go-gin-CRUD/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":5000")
}
