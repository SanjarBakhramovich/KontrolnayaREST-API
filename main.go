package main

import (
	"KONTROLNAYAREST-API/handlers"
	"KONTROLNAYAREST-API/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	memStore := store.NewMemStore()
	handler := handlers.NewHandler(memStore)

	r.POST("/books", handler.AddBook)
	r.GET("/books/:id", handler.GetBook)
	r.PUT("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)
	r.GET("/books", handler.GetAllBooks)

	r.Run() // По умолчанию на порту :8080
}
