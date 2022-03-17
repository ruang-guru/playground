package main

import (
	"fmt"
)

// Demo:
// Add Data Struct
// - Input: Name, Age, Job, Address, Add Again
// - Output: Data Users Updated

// Contoh:
// Input:
// - Add Name : Gina
// - Add Age : 30
// - Add Job : Designer
// - Add Address : Jakarta

// Output:
// - Users Added:  [{Aditira 20 Programmer Cianjur} {Dito 20 Programmer Bogor} {Gina 30 Designer Jakarta}]

// Input:
// - Ingin menambah user baru?(yes/no): if no (break) --> if yes (add again)

// Output:
// - Data Users Updated:
// - Name:  Aditira  Age:  20  Role:  Programmer  Address:  Cianjur
// - Name:  Dito  Age:  20  Role:  Programmer  Address:  Bogor
// - Name:  Gina  Age:  30  Role:  Designer  Address:  Jakarta
// - Name:  Tino  Age:  22  Role:  DevOps  Address:  Cina

type Person struct {
	name    string
	age     int
	job     string
	address string
}

func main() {
	var usersData = []Person{
		{
			name:    "Aditira",
			age:     20,
			job:     "Programmer",
			address: "Cianjur",
		},
		{
			name:    "Dito",
			age:     20,
			job:     "Programmer",
			address: "Bogor",
		},
	}

	for {
		fmt.Println("Data Users: ")
		for _, u := range usersData {
			fmt.Println("Name: ", u.name, " Age: ", u.age, " Role: ", u.job, " Address: ", u.address)
		}

		var name, job, address, addAgain string
		var age int
		fmt.Printf("Add Name : ")
		fmt.Scan(&name)
		fmt.Printf("Add Age : ")
		fmt.Scan(&age)
		fmt.Printf("Add Job : ")
		fmt.Scan(&job)
		fmt.Printf("Add Address : ")
		fmt.Scan(&address)

		var userStruct Person = Person{
			name:    name,
			age:     age,
			job:     job,
			address: address,
		}

		usersData = append(usersData, userStruct)
		fmt.Println("Users Added: ", usersData)

		fmt.Printf("Ingin memesan menu lain?(yes/no): ")
		fmt.Scan(&addAgain)

		if addAgain == "no" {
			fmt.Println("Data Users Updated: ")
			for _, u := range usersData {
				fmt.Println("Name: ", u.name, " Age: ", u.age, " Role: ", u.job, " Address: ", u.address)
			}

			break
		}

	}
}
