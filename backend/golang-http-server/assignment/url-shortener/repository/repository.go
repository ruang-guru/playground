package repository

import (
	"math/rand"
	"strings"
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	val, ok := r.Data[path]
	if !ok {
		return nil, entity.ErrURLNotFound
	}
	return &entity.URL{
		LongURL:  val,
		ShortURL: path,
	}, nil // TODO: replace this
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var shortUrl strings.Builder
	for {
		shortUrl.Reset()
		for i := 0; i < 6; i++ {
			shortUrl.Write([]byte(string(letters[rand.Int31n(int32(len(letters)))])))
		}
		_, ok := r.Data[shortUrl.String()]
		if ok {
			continue
		}
		break
	}

	r.Data[shortUrl.String()] = longURL

	return &entity.URL{
		LongURL:  longURL,
		ShortURL: shortUrl.String(),
	}, nil // TODO: replace this
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	_, ok := r.Data[customPath]
	if ok || customPath == "" {
		return nil, entity.ErrCustomURLIsExists
	}
	r.Data[customPath] = longURL

	return &entity.URL{LongURL: longURL, ShortURL: customPath}, nil // TODO: replace this
}
