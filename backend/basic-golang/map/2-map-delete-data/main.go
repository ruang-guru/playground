package main

import "fmt"

// Nah disini kita bisa menghapus data pada map menggunakan key nya.

func main() {
	var kamus = map[string]string{}
	kamus["sekolah"] = "bangunan atau lembaga untuk belajar dan mengajar serta tempat menerima dan memberi pelajaran"
	kamus["bisnis"] = "usaha komersial dalam dunia perdagangan; bidang usaha; usaha dagang"
	kamus["manajemen"] = "pimpinan yang bertanggung jawab atas jalannya perusahaan dan organisasi"

	// Disini kita akan menghapuskan data pada map dengan key "manajemen"
	for k, v := range kamus {
		fmt.Println(k, ":", v)
	}
	delete(kamus, "manajemen")
	fmt.Println()
	for k, v := range kamus {
		fmt.Println(k, ":", v)
	}
}
