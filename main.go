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
	c.JSON(http.StatusOK, books) ////its show true or false
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id { ///////if found it will  print
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func CreateBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil { /////  responsible for Binding  the JSON data to the newBook variable
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	index := -1
	for i, book := range books { ///// Find the index of the book with the matching ID
		if book.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found sorrtyyy"})
		return
	}

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil { ///////  responsible for Binding  the JSON data to the updatedBook variable
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) ////statusbadrequest  400 bad request
		return
	}

	// Update the book in the books slice
	books[index] = updatedBook

	c.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i, book := range books { // Find the index of the book with the matching ID
		if book.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	books = append(books[:index], books[index+1:]...) // Remove the book from the books slice

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func main() {
	router := gin.Default() //create a new router using gin framework

	books = append(books, Book{ID: "1", Title: "golang", Author: "rajen"})
	books = append(books, Book{ID: "2", Title: "gin", Author: "adarsh"})

	router.GET("/books", GetBooks)    ///when someone visit that url then the server will execute the /books
	router.GET("/books/:id", GetBook) ///getbook call the method
	router.POST("/books", CreateBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)

	router.Run(":8000") ////rounter.Run(localhost:8000)
}
