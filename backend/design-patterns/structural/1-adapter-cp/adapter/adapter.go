package adapter

import "fmt"

// Bayangin disini kita mempunyai AudioPlayer yang mempunyai method play dibawah ini.
// Method Play dari AudioPlayer ini menerima sebuah tipe Mp3
type AudioPlayer interface {
	Play(m Mp3) string
}

// Disini kita mempunyai Mp3 yang memiliki data.
// Anggap saja data ini adalah sebuah format yang bisa kita putar pada sebuah Mp3
type Mp3 struct {
	Data []byte
}

func (m Mp3) PlayAudio() string {
	fmt.Println("playing mp3 with data" + string(m.Data))
	return string(m.Data)
}

// anggaplah kita mempunya sebuah kaset yang mempunya PitaMusik dengan tipe string
type Kaset struct {
	PitaMusik string
}

// Walkman ini adalah sebuah AudioPLayer usang(legacy)
type Walkman struct{}

// Nah untuk memainkan sebuah PitaMusik. Kita putarkan Kaset tersebut pada walkman
func (w Walkman) Play(c Kaset) string {
	fmt.Println(c.PitaMusik)
	return c.PitaMusik
}

// Nah tapi bisa jika kita ingin memainkan sebuah Mp3 di sebuah Walkman ?
type Mp3ToKasetAdapter struct {
	Adaptee Walkman
}

func (a Mp3ToKasetAdapter) Play(m Mp3) string {
	return "" // TODO: replace this
}
