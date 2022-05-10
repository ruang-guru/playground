package main

import (
	"github.com/ruang-guru/playground/backend/design-patterns/creational/1-factory-pattern-cp/factoryLanjutan"
)

// Ini adalah file client, kita akan menggunakan package yang sudah kita buat sebelumnya
// Coba lah bermain main dengan kode ini dan gunakan contentCreator dengan type NetflixKorea
func main() {
	var contentCreator factoryLanjutan.ContentCreator
	contentCreator = &factoryLanjutan.NetflixKorea{}

	content := contentCreator.Produce(factoryLanjutan.Sunday)
	content.Play()
}
