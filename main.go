package main

import (
	"KONTROLNAYAREST-API/handlers"
	"KONTROLNAYAREST-API/store"

	"github.com/gin-gonic/gin"
)

// Обработчик для корневого URL
// func welcomePage(c *gin.Context) {
// 	c.HTML(http.StatusOK, "welcome.html", gin.H{
// 		"title": "Добро пожаловать!",
// 	})
// }

func main() {
	r := gin.Default()

	// r.GET("/", welcomePage)

	memStore := store.NewMemStore()
	handler := handlers.NewHandler(memStore)

	r.POST("/books", handler.AddBook)
	r.GET("/books/:id", handler.GetBook)
	r.PUT("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)
	r.GET("/books", handler.GetAllBooks)

	r.Run() // По умолчанию на порту :8080
}

// test Make
