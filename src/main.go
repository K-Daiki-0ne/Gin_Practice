package main

import (
	"0ne/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/", controllers.HelloWorld)

	app.Run(":3000")
}
