package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/design-patterns/structural/1-adapter-cp/adapter"
)

func main() {
	mp3 := adapter.Mp3{Data: []byte(`this is taylor swift song`)}
	walkman := adapter.Walkman{}

	ad := adapter.Mp3ToKasetAdapter{Adaptee: walkman}
	fmt.Println(ad.Play(mp3))
}
