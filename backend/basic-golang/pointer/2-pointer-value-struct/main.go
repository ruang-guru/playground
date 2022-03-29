package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//Method menggunakan pointer receiver
//Ketika mengubah suatu value disini, maka value dari objek juga akan berubah
func (p *Person) BirthdayPointer() {
	p.Age++
	fmt.Println("dalam method BirhtdayPointer umurnya", p.Age)
}

//Method menggunakan value receiver
//Ketika mengubah suatu value disini, maka value dari objek tidak ikut berubah
func (p Person) BirthdayValue() {
	p.Age++
	fmt.Println("dalam method BirthdayValue umurnya", p.Age)

}

func main() {
	var p Person
	p.Name = "putra"
	p.Age = 20
	fmt.Println("sebelum ulang tahun", p.Age)
	p.BirthdayPointer()
	fmt.Println("setelah memanggil BirthdayPointer", p.Age)

	p.BirthdayValue()
	fmt.Println("setelah memanggil BirthdayValue", p.Age)
}
