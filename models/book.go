package models

type Book struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	YearPublished    int    `json:"yearPublished"`
	Available        bool   `json:"available"`
	CoverImage       string `json:"coverImage"`
	FirstPublishYear int    `json:"firstPublishYear"`
}

type BookStore interface {
	AddBook(book *Book) error
	GetBook(id int) (*Book, error)
	UpdateBook(id int, book *Book) error
	DeleteBook(id int) error
	GetAllBooks() ([]*Book, error)
}
