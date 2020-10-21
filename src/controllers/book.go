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
	id := c.Param("id")
	singleBook := models.Single(id)
	c.JSON(http.StatusOK, singleBook)
}

// NewBook post new book controller
func NewBook(c *gin.Context) {
	title := c.Param("title")
	author := c.Param("author")
	rating := c.Param("rating")
	newBook := models.New(title, author, rating)
	c.JSON(http.StatusOK, newBook)
}

// DeleteBook delete single book controller
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	deleteBook := models.Delete(id)
	c.JSON(http.StatusOK, deleteBook)
}
