package main

import "fmt"

// Disini kita akan belajar tentang map
// map itu kalian bisa anggap sebagai kamus, terdapat kata:definisi
// atau dapat dikatakan key:value

func main() {

	// Contoh inisalisasi map
	/*

		map[type]type
		var m0 map[string]int{}
		m1 := make(map[string]int)
		m := map[string]int{"abc": 123, "xyz": 789} <- langsung di isi dengan key dan value

	*/

	// Kita pakai contoh seperti perumpaan di atas

	kamus := map[string]string{"sekolah": "bangunan atau lembaga untuk belajar dan mengajar serta tempat menerima dan memberi pelajaran",
		"bisnis":    "usaha komersial dalam dunia perdagangan; bidang usaha; usaha dagang",
		"manajemen": "pimpinan yang bertanggung jawab atas jalannya perusahaan dan organisasi"}

	// untuk mengakses definisi (value) kita harus mengetahui key nya

	fmt.Println(kamus["sekolah"])
	fmt.Println(kamus["bisnis"])
	fmt.Println(kamus["manajemen"])
	// Jika kita mencoba untuk melakukan print pada key yang tidak ada
	// Maka tidak ada yang ditampilkan
	fmt.Println(kamus["hilang"])

	// Kita juga dapat melakukan pengecekan apakah key tersebut ada atau tidak
	// val adalah value dari key yang kita cari
	// dan ok untuk melakukan pengecekan terhadap key nya apakah ada atau tidak
	// value dari ok bertipe data boolean (true atau false)
	if val, ok := kamus["sekolah"]; ok {
		fmt.Println("Key sekolah ada dan value nya adalah: ", val)
	} else {
		fmt.Println("Key sekolah tidak ada")
	}

	// Kalian masih ingat dengan `for range`. di map kita bisa melakukan itu
	for key, value := range kamus {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}
