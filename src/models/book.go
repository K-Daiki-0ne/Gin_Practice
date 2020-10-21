package models

import (
	"0ne/src/database"

	"github.com/jinzhu/gorm"
)

// Book book type
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

// All get all Book
func All() []Book {
	db := database.DBConnect
	var books []Book
	db.Find(&books)
	return books
}

// Single get single book
func Single(id string) Book {
	db := database.DBConnect
	var book Book
	db.Find(&book, id)
	return book
}

// New create new store book
func New(title string, author string, rating string) Book {
	db := database.DBConnect
	var newBook Book
	newBook.Title = title
	newBook.Author = author
	newBook.Rating = rating
	db.Create(&newBook)

	return newBook
}

// Delete delete book function
func Delete(id string) string {
	db := database.DBConnect
	var deleteBook Book
	var message string
	db.First(&deleteBook, id)
	if deleteBook.Title == "" {
		message = "No book found with given ID"
	} else {
		message = "Book Successfully deleted"
	}
	return message
}
