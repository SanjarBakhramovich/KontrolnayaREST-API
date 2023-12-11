package store

import (
	"KONTROLNAYAREST-API/models"
	"fmt"
	"sync"
)

// BookStore определяет интерфейс для работы с книгами
type BookStore interface {
	AddBook(book *models.Book) error
	GetBook(id int) (*models.Book, error)
	UpdateBook(id int, book *models.Book) error
	DeleteBook(id int) error
	GetAllBooks() ([]*models.Book, error)
}

// MemStore реализует интерфейс BookStore
type MemStore struct {
	sync.RWMutex
	books  map[int]*models.Book
	nextID int
}

// NewMemStore создает новый экземпляр MemStore
func NewMemStore() *MemStore {
	return &MemStore{
		books:  make(map[int]*models.Book),
		nextID: 1,
	}
}

// AddBook добавляет новую книгу в хранилище
func (m *MemStore) AddBook(book *models.Book) error {
	m.Lock()
	defer m.Unlock()
	book.ID = m.nextID
	m.books[m.nextID] = book
	m.nextID++
	return nil
}

// GetBook возвращает книгу по ее ID или ошибку, если книга не найдена
func (m *MemStore) GetBook(id int) (*models.Book, error) {
	m.RLock()
	defer m.RUnlock()

	book, ok := m.books[id]
	if !ok {
		return nil, fmt.Errorf("book with ID %d not found", id)
	}

	return book, nil
}

// UpdateBook обновляет информацию о книге по ее ID
func (m *MemStore) UpdateBook(id int, book *models.Book) error {
	m.Lock()
	defer m.Unlock()

	// Проверяем существование книги с указанным ID
	if _, ok := m.books[id]; !ok {
		return fmt.Errorf("book with ID %d not found", id)
	}

	// Обновляем информацию о книге
	m.books[id] = book

	return nil
}

// DeleteBook удаляет книгу по ее ID
func (m *MemStore) DeleteBook(id int) error {
	m.Lock()
	defer m.Unlock()

	// Проверяем существование книги с указанным ID
	if _, ok := m.books[id]; !ok {
		return fmt.Errorf("book with ID %d not found", id)
	}

	// Удаляем книгу
	delete(m.books, id)

	return nil
}

// GetAllBooks возвращает список всех книг
func (m *MemStore) GetAllBooks() ([]*models.Book, error) {
	m.RLock()
	defer m.RUnlock()

	var books []*models.Book
	for _, book := range m.books {
		books = append(books, book)
	}

	return books, nil
}
