package main

import (
	"bufio"
	"github.com/pkg/profile"
	"io"
	"log"
	"os"
	"testing"
	"unicode"
)

const fileName = "./lorem.txt"

var buf [1]byte

func readbyte(r io.Reader) (rune, error) {
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func mainOne() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open file %q: %v", fileName, err)
	}
	defer f.Close()

	words := 0
	inword := false

	for {
		r, err := readbyte(f)
		if err != nil {
			break
		}
		if err != nil {
			log.Fatalf("could not open file %q: %v", fileName, err)
		}

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
}

/*
dalam mainTwo() kita melakukan buffer pada NewReader.
karena kita tahu sebelumnya syscall.syscall() tidak melakukan
buffer pada file yang kita baca
*/
func mainTwo() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open file %q: %v", fileName, err)
	}
	defer f.Close()

	words := 0
	inword := false

	// The buffer to readbyte
	b := bufio.NewReader(f)

	for {
		r, err := readbyte(b)
		if err != nil {
			break
		}
		if err != nil {
			log.Fatalf("could not open file %q: %v", fileName, err)
		}

		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
}

func TestRunMainOne(t *testing.T) {
	mainOne()
}

// this function is used to only run MainTwo function
func TestRunMainTwo(t *testing.T) {
	mainTwo()
}

/*
go test -bench=BenchmarkMainOne -benchtime=1x -benchmem -cpuprofile profile.out
*/
func BenchmarkMainOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainOne()
	}
}

/*
go test -bench=BenchmarkMainTwo -benchtime=1x -benchmem -cpuprofile profile.out
*/
func BenchmarkMainTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainTwo()
	}
}
