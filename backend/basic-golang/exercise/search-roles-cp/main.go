package main

import (
	"fmt"
	"strings"
)

// Check Point:
// Search Roles
// - Input: Role
// - Output: Rearch Result

// Contoh:
// Data Users: [{Aditira 20 Programmer} {Dimas 18 Designer} {Yuni 21 DevOps} {Dito 22 Programmer} {Juno 25 DevOps}]

// Input:
//   - Masukan Role : Programmer

// Output:
// Programmer Found:
// Name:  Aditira  Age:  20  Role:  Programmer
// Name:  Dito  Age:  22  Role:  Programmer

// Input:
//   - Masukan Role : Secretary

// Output:
// Role: Sercretary Not Found!

type User struct {
	name string
	age  int
	role string
}

func main() {

	users := []User{
		{
			name: "Aditira",
			age:  20,
			role: "Programmer",
		},
		{
			name: "Dimas",
			age:  18,
			role: "Designer",
		},
		{
			name: "Yuni",
			age:  21,
			role: "DevOps",
		},
		{
			name: "Dito",
			age:  22,
			role: "Programmer",
		},
		{
			name: "Juno",
			age:  25,
			role: "DevOps",
		},
	}

	// TODO: answer here

	var (
		role   string
		result []int
	)

	fmt.Print("Masukkan Role : ")
	fmt.Scan(&role)

	for index, val := range users {
		if strings.EqualFold(strings.ToLower(val.role), strings.ToLower(role)) {
			result = append(result, index)
		}
	}

	if len(result) > 0 {
		fmt.Println(role, "Found")
		for _, val := range result {
			fmt.Printf("Name : %s \t Age : %d \t Role : %s\n", users[val].name, users[val].age, users[val].role)
		}
	} else {
		fmt.Printf("Role : %s Not Found!", role)
	}
}
