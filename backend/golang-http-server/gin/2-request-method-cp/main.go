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
	route := gin.Default()
	route.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	route.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})

	route.PUT("/put", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})

	route.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	route.PATCH("/patch", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "PATCH",
		})
	})

	route.HEAD("/head", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "HEAD",
		})
	})

	route.OPTIONS("/options", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OPTIONS",
		})
	})

	return route // TODO: replace this
}

func main() {
	r := GetGinRoute()
	r.Run("localhost:3000")
}
