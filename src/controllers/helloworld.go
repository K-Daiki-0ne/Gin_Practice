package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloWorld controller
func HelloWorld(c *gin.Context) {
	response := "Hello World"
	c.JSON(http.StatusOK, response)
}
