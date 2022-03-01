package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

var movies []Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}
	result, err := json.Marshal(movies)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}
	var newMovie Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error %s", err)
		return
	}
	newMovie.Id = len(movies) + 1
	log.Printf("added new movie : %+v", newMovie)
	movies = append(movies, newMovie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	key := vars["id"]
	intKey, err := strconv.Atoi(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error %s", err)
		return
	}
	var result []byte
	id, available := findFilm(intKey)
	if available {
		result, err = json.Marshal(movies[id])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("error %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("not found")
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func findFilm(key int) (int, bool) {
	for i := 0; i < len(movies); i++ {
		if key == movies[i].Id {

			return i, true
		}
	}
	return -1, false
}
