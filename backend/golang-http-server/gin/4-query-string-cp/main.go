package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	// TODO: answer here
	router.GET("/movie", func(ctx *gin.Context) {
		genre := ctx.DefaultQuery("genre", "general")
		country := ctx.Query("country")

		var res strings.Builder
		res.WriteString(fmt.Sprintf("Here result of query of movie with genre %s", genre))

		if country != "" {
			res.WriteString(fmt.Sprintf(" in country %s", country))
		}

		ctx.String(http.StatusOK, res.String())
	})
	return router
}

func main() {
	router := GetRouter()
	router.Run(":8080")
}
