package main

import (
	"0ne/src/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Router(app)
	app.Run(":3000")
}
