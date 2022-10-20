package repository

import (
	"errors"
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

func (repo *Repository) GetBook(id string) (domain.Book, error) {
	if book, exist := repo.DB[id]; exist {
		return book, nil
	}

	return domain.Book{}, errors.New("not found")
}

func (repo *Repository) CreateBook(book domain.Book) (domain.Book, error) {
	if _, exist := repo.DB[book.ID]; exist {
		return domain.Book{}, errors.New("already exist")
	}

	repo.DB[book.ID] = book

	return book, nil
}

func (repo *Repository) UpdateBook(book domain.Book) error {
	if _, exist := repo.DB[book.ID]; !exist {
		return errors.New("not found")
	}

	repo.DB[book.ID] = book

	return nil
}

func (repo *Repository) DeleteBook(id string) error {
	if _, exist := repo.DB[id]; !exist {
		return errors.New("not found")
	}

	delete(repo.DB, id)

	return nil
}
