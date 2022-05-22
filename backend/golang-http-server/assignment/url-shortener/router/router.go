package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/handlers"
)

func SetupRouter(urlHandler handlers.URLHandler) *gin.Engine {
	engine := gin.Default()
	engine.POST("/custom", func(ctx *gin.Context) {
		urlHandler.CreateCustom(ctx)
	})
	engine.POST("/", func(ctx *gin.Context) {
		urlHandler.Create(ctx)
	})
	engine.GET("/:short", func(ctx *gin.Context) {
		urlHandler.Get(ctx)
	})
	return engine // TODO: replace this
}
