package api

import (
	"encoding/json"
	"net/http"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: ""}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this
}
