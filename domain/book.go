package domain

type Book struct {
	ID     string `json:"id"`
	Tittle string `json:"title" binding:"required"`
	Year   int    `json:"year" binding:"required"`
	Author string `json:"author" binding:"required"`
}
