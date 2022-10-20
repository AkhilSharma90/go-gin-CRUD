package services

import (
	"go-gin-CRUD/domain"
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

func (s *Services) GetAllBooks() ([]domain.Book, error) {
	return s.Repository.GetAllBooks()
}

func (s *Services) GetBook(id string) (domain.Book, error) {
	return s.Repository.GetBook(id)
}
