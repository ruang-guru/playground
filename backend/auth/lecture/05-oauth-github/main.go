package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const clientID = "a343d4093b482690cf50"
const clientSecret = "2ce58511e2bfd92dc707b5d3296b10d3752d2991"

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Buat redirect url for github.com
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		// Pada handler ini, ambil value code yang dikirim sebagai code auth dari github
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		// Call github oauth endpoint token menggunakan client id dan
		// client secret serta code yang telah dikirimkan ketika redirect auth github
		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		// Set request accept json
		req.Header.Set("accept", "application/json")

		//Send the http client request
		httpClient := http.Client{}
		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer res.Body.Close()

		// Parse response to OauthAccessResponse
		var t OAuthAccessResponse
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		err = json.Unmarshal(body, &t)
		if err != nil {
			log.Fatal(readErr)
		}

		// Redirect response to welcome.html dengan access token
		fmt.Println(t.AccessToken)
		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})

	fmt.Println("Starting Server at port :8000")
	http.ListenAndServe(":8000", nil)
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
