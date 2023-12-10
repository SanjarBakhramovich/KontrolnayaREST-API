package handlers

import (
	"KONTROLNAYAREST-API/models"
	"KONTROLNAYAREST-API/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handler структура, включающая в себя зависимости для обработчиков
type Handler struct {
	Store store.BookStore
}

// NewHandler создает новый экземпляр Handler с зависимостями
func NewHandler(s store.BookStore) *Handler {
	return &Handler{Store: s}
}

// GetAllBooks обрабатывает GET запрос для получения списка всех книг
func (h *Handler) GetAllBooks(c *gin.Context) {
	books, err := h.Store.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook обрабатывает GET запрос для одной книги
func (h *Handler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	book, err := h.Store.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func extractAuthors(authorsData interface{}) string {
	authorsSlice, ok := authorsData.([]interface{})
	if !ok {
		return "" // или обработайте ошибку по-другому
	}

	var authors []string
	for _, author := range authorsSlice {
		authorMap, ok := author.(map[string]interface{})
		if !ok {
			continue // или обработайте ошибку по-другому
		}
		if name, ok := authorMap["name"].(string); ok {
			authors = append(authors, name)
		}
	}

	return strings.Join(authors, ", ")
}

// fetchBookInfoFromOpenLibrary делает запрос к Open Library API и возвращает информацию о книге
func (h *Handler) fetchBookInfoFromOpenLibrary(isbn string) (*models.Book, error) {
	url := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&format=json&jscmd=data", isbn)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Получаем данные по ключу ISBN
	bookData, ok := result["ISBN:"+isbn].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("no data found for ISBN %s", isbn)
	}

	// Создаем экземпляр Book и заполняем его данными
	book := models.Book{
		Title: bookData["title"].(string),
		// Пример для автора, предполагая, что авторы возвращаются в виде массива
		Author: extractAuthors(bookData["authors"]),
		// Аналогично для других полей...
	}

	return &book, nil
}

// AddBook обрабатывает POST запрос для добавления новой книги
func (h *Handler) AddBook(c *gin.Context) {
	var input struct {
		ISBN string `json:"isbn"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.fetchBookInfoFromOpenLibrary(input.ISBN)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.AddBook(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook обрабатывает PUT запрос для обновления книги
func (h *Handler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Store.UpdateBook(id, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook обрабатывает DELETE запрос для удаления книги
func (h *Handler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.Store.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting book"})
		return
	}

	c.Status(http.StatusNoContent)
}
