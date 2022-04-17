package main

import (
	"fmt"
	"time"
)

func main() {
	imageQueue := make(chan string, 100)
	imageResizer := func() {
		//infinite loop
		for {
			image := <-imageQueue
			fmt.Printf("%s resized\n", image)
			// do something
		}
	}
	go imageResizer()

	//common business logic
	//when we need to resize image, we can just put it into the queue
	imageQueue <- "image.jpg"
	imageQueue <- "image2.jpg"
	time.Sleep(100 * time.Millisecond)
}
