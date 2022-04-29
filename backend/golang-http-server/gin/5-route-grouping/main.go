package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login v1")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submit v1")
		})
		v1.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "read v1")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login v2")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "submit v2")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "read v2")
		})
	}

	router.Run(":8080")
}
