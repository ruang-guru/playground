package main

// Dalam golang, untuk membuat sebuah client HTTP kita dapat menggunakan package net/http.
// Pada package net/http terdapat struct Client yang dapat digunakan untuk merepresentasikan client HTTP.

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Membuat client http
	c := http.Client{}

	// Melakukan request ke https://animechan.vercel.app/api/random dengan method GET
	resp, err := c.Get("https://animechan.vercel.app/api/random")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer resp.Body.Close()

	// Membaca response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf(string(body))
}
