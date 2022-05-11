package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here

	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
