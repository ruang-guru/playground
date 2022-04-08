package main

// Dari contoh yang telah diberikan, cobalah untuk membuat stack yang memiliki jumlah maksimal elemen sebanyak 10 dengan menambahkan atribut Size.
// Gunakan slice untuk menyimpan data pada stack. Pada kasus ini, data yang disimpan berupa int.
// Lengkapi function NewStack sehingga function tersebut dapat mengembalikan objek Stack dengan memiliki ukuran dari parameter.
type Stack struct {
	//beginanswer
	Top  int
	Size int
	Data []int
	//endanswer
}

func NewStack(size int) Stack {
	//beginanswer
	return Stack{
		Top:  -1,
		Size: size,
		Data: []int{},
	}
	//endanswer
}
