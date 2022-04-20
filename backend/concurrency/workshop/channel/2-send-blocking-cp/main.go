package main

import "fmt"

func sendBlock(output chan bool) {
	c := make(chan bool)
	called := false
	go func() {
		fmt.Println("receive from main")
		//memberi called nilai dari channel c
		// TODO: answer here
		called = <-c
	}()

	//mengirim bool value true ke channel c
	// TODO: answer here
	c <- true
	output <- called
	fmt.Println(c) //agar variabel c digunakan
}
