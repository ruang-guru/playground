package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type message struct {
	Data string `json:"data"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var newUser user
		json.Unmarshal(reqBody, &newUser)
		log.Printf("i receive this data: %v\n", newUser)
		var msg message
		msg.Data = fmt.Sprintf("hello %s, your age is %d", newUser.Name, newUser.Age)
		log.Printf("i send this data: %v\n", msg)
		resultByte, _ := json.Marshal(msg)
		w.WriteHeader(http.StatusAccepted)
		w.Write(resultByte)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	port := ":8010"
	http.HandleFunc("/hello", hello) //saat mengakses route /hello yang dijalankan adalah fungsi hello
	log.Printf("server started on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

//referensi: https://dasarpemrogramangolang.novalagung.com/A-web-service-api.html
