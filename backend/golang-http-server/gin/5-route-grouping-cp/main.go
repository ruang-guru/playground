package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{}) // TODO: replace this
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
