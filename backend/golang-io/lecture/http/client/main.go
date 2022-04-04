package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type message struct {
	Data string `json:"data"`
}

var baseURL = "http://localhost:8010"

func checkServer(user user) (message, error) {
	var err error
	var client = &http.Client{}
	userByte, _ := json.Marshal(user)
	var receivedMessage message
	request, err := http.NewRequest("POST", baseURL+"/hello", bytes.NewBuffer(userByte))
	if err != nil {
		return message{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return message{}, err
	}
	defer response.Body.Close()

	rBody, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(rBody, &receivedMessage)
	if err != nil {
		return message{}, err
	}

	return receivedMessage, nil
}

func main() {
	var name string
	var age int
	fmt.Println("enter your name:")
	fmt.Scan(&name)
	fmt.Println("enter your age:")
	fmt.Scanf("%d", &age)
	//line diatas hanya untuk mengambil input kita
	user1 := user{
		Name: name,
		Age:  age,
	}
	fmt.Printf("i send this data: %v\n", user1)
	receivedMessage, _ := checkServer(user1)
	fmt.Printf("i receive this data: %v\n", receivedMessage)
}
