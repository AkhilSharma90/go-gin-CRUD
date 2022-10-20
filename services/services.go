package services

import (
	"go-gin-CRUD/domain"
	"go-gin-CRUD/repository"

	"github.com/google/uuid"
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

func (s *Services) CreateBook(book domain.Book) (domain.Book, int, error) {
	book.ID = uuid.New().String()

	book, err := s.Repository.CreateBook(book)
	if err != nil {
		return domain.Book{}, 400, err
	}

	return book, 201, nil
}

func (s *Services) UpdateBook(book domain.Book) (domain.Book, error) {
	if err := s.Repository.UpdateBook(book); err != nil {
		return domain.Book{}, err
	}

	return book, nil
}
