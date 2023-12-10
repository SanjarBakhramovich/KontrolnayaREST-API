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

	r.GET("/books", handler.GetAllBooks)
	r.GET("/books/:id", handler.GetBook)
	r.POST("/books", handler.AddBook)
	r.PUT("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)

	r.Run() // По умолчанию на порту :8080
}
