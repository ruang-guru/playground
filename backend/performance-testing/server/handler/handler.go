package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruang-guru/playground/backend/performance-testing/server/entity"
	"github.com/ruang-guru/playground/backend/performance-testing/server/repository"
)

type Handler interface {
	GetMovies(c *gin.Context)
	GetMovie(c *gin.Context)
	AddMovie(c *gin.Context)
	Home(c *gin.Context)
}

type handler struct {
	movies repository.Repository
}

func New(repo repository.Repository) Handler {
	return &handler{repo}
}

func (h *handler) GetMovies(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, fmt.Errorf("invalid request method"))
		return
	}
	result, err := h.movies.GetAll()
	resultSlice := []entity.Movie{}
	for _, movie := range result {
		resultSlice = append(resultSlice, movie)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		log.Printf("error %s", err)
		return
	}
	c.JSON(http.StatusOK, resultSlice)
}

func (h *handler) AddMovie(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, fmt.Errorf("invalid request method"))
		return
	}
	var newMovie entity.Movie
	err := c.BindJSON(&newMovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		log.Printf("error %s", err)
		return
	}
	err = h.movies.Store(&newMovie)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		log.Printf("error %s", err)
		return
	}
	log.Printf("added new movie : %+v", newMovie)
	c.JSON(http.StatusOK, "new movie added")
}

func (h *handler) GetMovie(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, fmt.Errorf("invalid request method"))
		return
	}
	key := c.Param("id")
	intKey, err := strconv.Atoi(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		log.Printf("error %s", err)
		return
	}
	result, err := h.movies.Get(intKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		log.Printf("error %s", err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *handler) Home(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, "bad request")
			return
		}
		log.Print(string(data))
	}
	c.Writer.WriteHeader(http.StatusOK)
}
