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
	paramQuery := req.URL.Query()
	username, err := api.usersRepo.Login(paramQuery.Get("username"), paramQuery.Get("password"))
	if err != nil {
		errorResp := AuthErrorResponse{
			Error: err.Error(),
		}
		errorJson, _ := json.Marshal(errorResp)
		http.Error(w, string(errorJson), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: *username}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	_, err := api.usersRepo.FindLoggedinUser()
	if err != nil {
		errorResp := AuthErrorResponse{
			Error: err.Error(),
		}
		errorJson, _ := json.Marshal(errorResp)
		http.Error(w, string(errorJson), http.StatusUnauthorized)
		return
	}
	paramQuery := req.URL.Query()
	err = api.usersRepo.Logout(paramQuery.Get("username"))

	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()}) // TODO: replace this
	}

}
