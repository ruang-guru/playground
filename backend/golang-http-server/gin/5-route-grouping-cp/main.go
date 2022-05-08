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
	c.JSON(http.StatusOK, gin.H{
		"data": movies,
	}) // TODO: replace this
}

var MovieGetHandler = func(c *gin.Context) {
	idMovie := c.Param("id")
	idInt, err := strconv.Atoi(idMovie)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	val, ok := movies[idInt]
	if !ok {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": val,
	}) // TODO: replace this
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	movies := router.Group("movie")
	{
		movies.GET("get/:id", MovieGetHandler)
		movies.GET("list", MovieListHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
