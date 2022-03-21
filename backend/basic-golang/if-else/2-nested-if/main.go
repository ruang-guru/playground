package main

import "fmt"

// Disin kita akan belajar tentang nested if
// if di dalam if

// studi kasus jika
func main() {
	nilai := 999
	if nilai > 100 {
		if nilai < 500 {
			fmt.Println("nilai diantara 100 dan 500")
		}
		if nilai > 500 {
			fmt.Println("nilai lebih dari 500")
			if nilai < 1000 {
				fmt.Println("nilai lebih dari 500 dan kurang dari 1000")
			}
		}
	}

}
