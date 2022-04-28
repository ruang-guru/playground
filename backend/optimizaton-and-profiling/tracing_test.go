package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/pkg/profile"
)

type Counts map[string]int

func TestMainProgramm(t *testing.T) {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	size := 10
	q := make(chan string, size)
	o := make(chan Counts, size)
	for i := 0; i < size*2; i++ {
		go contentAsync("https://www.ruangguru.com", q)
	}

	for i := 0; i < 5; i++ {
		go countAsync(q, o)
	}

	for i := 0; i < size*2; i++ {
		fmt.Println(len(<-o))
	}
}

func contentAsync(url string, output chan<- string) error {
	c, err := content(url)
	if err != nil {
		return err
	}
	output <- c
	return nil
}

func countWords(content string) Counts {
	out := Counts{}
	words := strings.Split(content, ",")
	for _, w := range words {
		out[w]++
	}
	return out
}

func countAsync(input <-chan string, output chan<- Counts) {
	for c := range input {
		output <- countWords(c)
	}
	close(output)
}

func content(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	out := &strings.Builder{}
	buff := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buff)
		if n <= 0 {
			break
		}
		out.Write(buff[:n])
	}
	// lets say there is delay about 200ms
	time.Sleep(200 * time.Millisecond)
	return out.String(), nil
}
