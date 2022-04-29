package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Di Gin, untuk mendapatkan path parameter yaitu kita memberikan tanda : pada route.
// Untuk mendapatkan nilai path parameter, kita dapat menggunakan c.Param() dengan nilai parameter adalah route yang memiliki tanda :

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/user/:name/:country", func(c *gin.Context) {
		name := c.Param("name")
		country := c.Param("country")
		c.String(http.StatusOK, "Hello %s from %s", name, country)
	})

	router.Run(":8080")
}
