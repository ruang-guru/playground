package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	http.HandleFunc("/registrations", registrationsHandler)
	http.ListenAndServe(":8081", nil)
}

func registrationsHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	if req.FormValue("username") == "" || req.FormValue("password") == "" {
		fmt.Fprintf(w, "Please enter a valid username and password.\r\n")
	} else {
		response, err := registerUser(req.FormValue("username"), req.FormValue("password"))
		if err != nil {
			fmt.Fprintf(w, "Error: %s\r\n", err)
		} else {
			fmt.Fprintf(w, "Successfully registered user.\r\n")
			fmt.Fprintf(w, "Hashed password: %s\r\n", response)
		}
	}
}

func registerUser(username string, password string) (string, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPassword), nil
}
