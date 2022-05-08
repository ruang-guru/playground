package main

import "github.com/gin-gonic/gin"

// Berbeda dengan standard library di Golang, untuk menghadle request method pada gin, kita dapat menggunakan
// method GET, POST, PUT, dan yang lainnya dari objek gin.Default.

func main() {
	router := gin.Default()

	// Request method GET
	router.GET("/someGet", func(c *gin.Context) {})

	// Request method POST
	router.POST("/somePost", func(c *gin.Context) {})

	// Request method PUT
	router.PUT("/somePut", func(c *gin.Context) {})

	// Request method DELETE
	router.DELETE("/someDelete", func(c *gin.Context) {})

	// Request method PATCH
	router.PATCH("/somePatch", func(c *gin.Context) {})

	// Request method HEAD
	router.HEAD("/someHead", func(c *gin.Context) {})

	// Request method OPTIONS
	router.OPTIONS("/someOptions", func(c *gin.Context) {})

	router.Run()
}
