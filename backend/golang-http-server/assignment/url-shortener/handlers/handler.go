package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/repository"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) URLHandler {
	return URLHandler{
		repo: repo,
	}
}

func (h *URLHandler) Get(c *gin.Context) {
	// TODO: answer here
	entityUrl, err := h.repo.Get(c.Param("short"))
	if err != nil {
		if err == entity.ErrURLNotFound {
			c.String(http.StatusNotFound, "not found")
			return
		}
		c.String(http.StatusInternalServerError, "error")
		return
	}
	c.Writer.WriteHeader(http.StatusFound)
	c.Header("Location", entityUrl.LongURL)
}

func (h *URLHandler) Create(c *gin.Context) {
	// TODO: answer here
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	urlEntity := entity.URL{}
	json.Unmarshal(reqBody, &urlEntity)
	respEntity, err := h.repo.Create(urlEntity.LongURL)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, struct {
		Data entity.URL
	}{
		Data: *respEntity,
	})
}

func (h *URLHandler) CreateCustom(c *gin.Context) {
	// TODO: answer here
	reqBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	urlEntity := entity.URL{}
	json.Unmarshal(reqBody, &urlEntity)
	respEntity, err := h.repo.CreateCustom(urlEntity.LongURL, urlEntity.ShortURL)
	if err != nil {
		if err == entity.ErrCustomURLIsExists {
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, struct {
		Data entity.URL
	}{
		Data: *respEntity,
	})
}
