package main

import (
	"github.com/AkhilSharma90/go-gin-CRUD/handlers"
	"github.com/AkhilSharma90/go-gin-CRUD/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":5000")
}
