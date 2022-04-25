package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	for {
		fmt.Printf("\x1bc") // clear screen
		var username, password, loginAgain string
		fmt.Println()
		fmt.Println("Basic Auth Login Form:")
		fmt.Printf("Username : ")
		fmt.Scan(&username)
		fmt.Printf("Password : ")
		fmt.Scan(&password)

		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/login", nil)
		if err != nil {
			panic(err)
		}

		req.SetBasicAuth(username, password)
		req.Header.Add("Content-Type", "application/json")
		req.Close = true

		client := http.Client{}
		response, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		if response.StatusCode != http.StatusOK {
			fmt.Println("non 2xx response from server, request: ", response.Status)

			fmt.Printf("Loggin Again?(y/n): ")
			fmt.Scan(&loginAgain)
			if loginAgain != "y" {
				fmt.Println("Thank you!")
				return

			}

		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		log.Print(string(body))

		fmt.Printf("Loggin Again?(y/n): ")
		fmt.Scan(&loginAgain)
		if loginAgain != "y" {
			fmt.Println("Thank you!")
			return

		}
	}
}
