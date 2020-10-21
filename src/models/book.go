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
