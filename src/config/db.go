package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// MySQL Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Book book type
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

const (
	// Dialect Use DataBase name
	Dialect = "mysql"

	// DBUser MySQL's user name
	DBUser = "user1"

	// DBPass MySQL's user password
	DBPass = "Password_01"

	// DBProt MySQL protocol
	DBProt = "tcp(localhost:3306)"

	// DBName MySQL table name
	DBName = "crud"
)

// Connect MySQL connect function
func Connect() *gorm.DB {
	connect := Parser(DBUser, DBPass, DBProt, DBName)

	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		fmt.Println("DB connect ...NO")
	} else {
		db.AutoMigrate(&Book{})
		fmt.Println("DB connect ...OK")
	}
	return db
}

// Close Database close function
func Close() {
	db := Connect()
	db.Close()
	fmt.Println("DB close ...OK")
}
