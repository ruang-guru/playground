package main

import (
	"fmt"
)

// Demo:
// Add Data Map
// - Input: Username, Email, Add Again
// - Output: Data Users Updated

// Contoh:
// Input:
// - Form:
//   - Masukan username baru: dita
//   - Masukan email: dita@gmail.com

// Output:
// - Data updated! dengan uername: dita dan email: dita@gmail.com

// Input:
// - Ingin menambah data baru?(yes/no): if no (break) --> if yes (add again)

// Output:
// Data Users Updated!:
// - Name: aditira ,  Email : aditira@gmail.com
// - Name: dito ,  Email : dito@gmail.com
// - Name: dita ,  Email : dita@gmail.com

func main() {
	users := map[string]string{
		"aditira": "aditira@gmail.com",
		"dito":    "dito@gmail.com",
	}

	for {
		fmt.Println("Data Users:")
		for name, email := range users {
			fmt.Println("- Name:", name, ", ", "Email :", email)
		}
		fmt.Println()
		var username, email, addAgain string

		fmt.Println("Form:")
		fmt.Printf("Masukan username baru: ")
		fmt.Scan(&username)
		fmt.Printf("Masukan email: ")
		fmt.Scan(&email)

		users[username] = email

		fmt.Println()
		fmt.Printf("Data updated! dengan username: %v dan email: %v\n\n", username, email)

		fmt.Printf("Ingin menambah data baru?(yes/no): ")
		fmt.Scan(&addAgain)

		if addAgain == "no" {
			break
		}
	}
	fmt.Println("Data Users Updated!:")
	for name, email := range users {
		fmt.Println("- Name:", name, ", ", "Email :", email)
	}
}
