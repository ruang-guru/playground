package graderapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ruang-guru/playground/cli/assignmentmanager/entity"
)

type LoginResponse struct {
	Data    entity.TokenData `json:"data"`
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Detail  string           `json:"detail"`
}

type Email struct {
	Email string `json:"email"`
}

func (g *GraderApiClient) SetTokenData(tokenData *entity.TokenData) {
	g.AccessToken = tokenData.AccessToken
	g.RefreshToken = tokenData.RefreshToken
}

func (g *GraderApiClient) Login() (*entity.TokenData, error) {
	postBody, _ := json.Marshal(map[string]string{
		"email":    g.Email,
		"password": g.Password,
	})
	reqBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(GraderApiUrl+"/user/login", "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var loginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to parse response body\n%s\n\n%s", string(responseBody), err.Error())
	}

	if loginResponse.Status != "success" {
		return nil, fmt.Errorf("[%s] %s", loginResponse.Message, loginResponse.Detail)
	}

	return &loginResponse.Data, nil
}

func (g *GraderApiClient) RegenerateAccessToken() (*entity.TokenData, error) {
	postBody, _ := json.Marshal(map[string]string{
		"refreshToken": g.RefreshToken,
	})
	reqBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(GraderApiUrl+"/user/refresh-token", "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var loginResponse LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to parse response body\n%s\n\n%s", string(responseBody), err.Error())
	}

	if loginResponse.Status != "success" {
		return nil, fmt.Errorf("[%s] %s", loginResponse.Message, loginResponse.Detail)
	}

	return &loginResponse.Data, nil
}
