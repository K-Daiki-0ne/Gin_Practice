package database

import (
	"github.com/jinzhu/gorm"
	// MySQL Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DBConnect config/db.go instance object
	DBConnect *gorm.DB
)
