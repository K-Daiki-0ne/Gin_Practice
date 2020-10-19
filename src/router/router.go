package router

import (
	"0ne/src/controllers"

	"github.com/gin-gonic/gin"
)

// Router this app router function
func Router(app *gin.Engine) {
	app.GET("/api/v1/books", controllers.GetAllBooks)
	app.GET("/api/v1/book/:id", controllers.GetBook)
	app.POST("/api/v1/new/book", controllers.NewBook)
	app.DELETE("/api/v1/delete/book/:id", controllers.DeleteBook)
}
