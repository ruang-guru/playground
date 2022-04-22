package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	username = "aditira"
	password = "aditira123"
)

func main() {
	callAPI("https://localhost:8080/login", "POST")
}

func callAPI(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}
	req.SetBasicAuth(username, password)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}
