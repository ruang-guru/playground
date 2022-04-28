package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	handler "github.com/ruang-guru/playground/backend/performance-testing/server/handler"
)

type api struct {
	engine  *gin.Engine
	handler handler.Handler
}

type API interface {
	InitAPI()
	InitPromeHandler()
}

func New(e *gin.Engine, handler handler.Handler) API {
	return &api{
		engine:  e,
		handler: handler,
	}
}

func (r *api) InitAPI() {
	routerGroup := r.engine.Group("/")
	routerGroup.GET("/movies", r.handler.GetMovies)
	routerGroup.GET("/movie/:id", r.handler.GetMovie)
	routerGroup.POST("/movie", r.handler.AddMovie)
	routerGroup.GET("/", r.handler.Home)
	routerGroup.POST("/", r.handler.Home)
}

func (r *api) InitPromeHandler() {
	promHandler := promhttp.Handler()
	r.engine.GET("/metrics", func(c *gin.Context) {
		promHandler.ServeHTTP(c.Writer, c.Request)
	})
}
