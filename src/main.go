package main

import (
	"0ne/src/config"
	"0ne/src/router"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Connect()
}

func main() {
	app := gin.Default()
	router.Router(app)
	app.Run(":3000")
	defer config.Close()
}
