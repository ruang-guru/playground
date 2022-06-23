package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var RequestMethodHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

func GetGinRoute() *gin.Engine {
	return &gin.Engine{} // TODO: replace this
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
