package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var (
	// Dialect Use DataBase name
	Dialect = "mysql"

	// DBUser MySQL's user name
	DBUser = os.Getenv("DB_USER")

	// DBPass MySQL's user password
	DBPass = os.Getenv("DB_PASS")

	// DBProt MySQL protocol
	DBProt = os.Getenv("DB_PROT")

	// DBName MySQL table name
	DBName = os.Getenv("DB_NAME")
)

// Connect MySQL connect function
func Connect() *gorm.DB {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProt, DBName)
	db, err := gorm.Open(Dialect, connect)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("DB connect ...OK")

	return db
}

// Close Database close function
func Close() {
	db := Connect()
	defer db.Close()
	fmt.Println("DB close ...OK")
}
