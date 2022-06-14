package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Salah satu kemudahan lain dari Gin adalah ketika kita ingin mengambil nilai query parameter.
// Untuk mengambil nilai query parameter kita dapat menggunakan method c.Query() dengan parameter berupa key dari query param.
// Kita juga dapat mengatur nilai default jika suatu key query param tidak ada dengan menggunakan c.DefaultQuery().

func main() {
	router := gin.Default()

	router.GET("/profile", func(c *gin.Context) {
		role := c.DefaultQuery("role", "guest")
		name := c.Query("name")

		c.String(http.StatusOK, "role: %s, name: %s", role, name)
	})
	router.Run(":8080")
}
