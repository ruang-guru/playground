package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	res, err := api.usersRepo.Login(username, password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	fmt.Println(res)

	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: ""}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	err := api.usersRepo.Logout(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this
}
