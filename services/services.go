package services

import (
	"go-gin-CRUD/repository"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices() *Services {
	return &Services{
		Repository: repository.NewRepository(),
	}
}
