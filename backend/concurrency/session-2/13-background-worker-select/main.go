package main

import "fmt"

func main() {

	imageQueue := make(chan string, 100)
	videoQueue := make(chan string, 100)
	quit := make(chan bool)
	resizer := func() {
		for {
			select {
			case image := <-imageQueue:
				fmt.Printf("%s resized\n", image)
			case video := <-videoQueue:
				fmt.Printf("%s resized\n", video)
			case <-quit:
				break
			}
		}
	}
	go resizer()

	imageQueue <- "image.jpg"
	videoQueue <- "video.mp4"
	quit <- true
}
