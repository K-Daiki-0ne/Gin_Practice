package controllers

import (
	"0ne/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllBooks all book get controller
func GetAllBooks(c *gin.Context) {
	allBooks := models.All()
	c.JSON(http.StatusOK, allBooks)
}

// GetBook single book get controller
func GetBook(c *gin.Context) {
	book := "Single Book"
	c.JSON(http.StatusOK, book)
}

// NewBook post new book controller
func NewBook(c *gin.Context) {
	newBook := "New Book"
	c.JSON(http.StatusOK, newBook)
}

// DeleteBook delete single book controller
func DeleteBook(c *gin.Context) {
	deleteBook := "Delete Book"
	c.JSON(http.StatusOK, deleteBook)
}
