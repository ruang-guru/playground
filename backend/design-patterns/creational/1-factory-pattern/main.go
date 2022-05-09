package main

import (
	"github.com/ruang-guru/playground/backend/design-patterns/creational/1-factory-pattern/factoryAwal"
)

// Ini adalah file client, kita akan menggunakan package yang sudah kita buat sebelumnya
// Coba lah bermain main dengan kode ini dan gunakan contentCreator dengan type NetflixKorea
func main() {
	var contentCreator factoryAwal.ContentCreator
	contentCreator = &factoryAwal.NetflixKorea{}
	content := contentCreator.Produce()
	content.Play()
}
