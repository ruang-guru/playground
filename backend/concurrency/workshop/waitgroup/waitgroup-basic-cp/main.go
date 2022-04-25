package main

import "sync"

func testWG(output chan<- []bool) {
	// TODO: answer here

	done := make([]bool, 5)

	for i := 0; i < 5; i++ {
		// TODO: answer here
		go func(i int) {
			// TODO: answer here
			done[i] = true
		}(i)
	}

	// TODO: answer here
	output <- done
}
