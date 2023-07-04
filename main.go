package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, book := range books {
		if book.ID == id {
			books[index] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted sucessfully "})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
	router := gin.Default()

	books = append(books, Book{ID: "1", Title: "golang", Author: "rajen"})
	books = append(books, Book{ID: "2", Title: "gin", Author: "adarsh"})

	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBook)
	router.POST("/books", CreateBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)

	router.Run(":8000")
}
