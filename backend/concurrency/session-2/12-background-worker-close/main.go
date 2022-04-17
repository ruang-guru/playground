package main

import (
	"fmt"
)

func main() {
	imageQueue := make(chan string, 100)
	c := make(chan int)
	imageResizer := func() {
		for image := range imageQueue {
			fmt.Printf("%s resized\n", image)
		}
		c <- 1
	}
	go imageResizer()

	imageQueue <- "image.jpg"
	imageQueue <- "image2.jpg"
	close(imageQueue)
	<-c
}
