package repository

import (
	"go-gin-CRUD/domain"
)

type Repository struct {
	DB map[string]domain.Book
}

func NewRepository() *Repository {
	example := domain.Book{
		ID:     "1",
		Tittle: "Example Tittle",
		Year:   2019,
		Author: "Example Author",
	}

	return &Repository{
		DB: map[string]domain.Book{example.ID: example},
	}
}

func (repo *Repository) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book

	for _, book := range repo.DB {
		books = append(books, book)
	}

	return books, nil
}
