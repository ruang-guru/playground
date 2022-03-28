package main

import "fmt"

// Interface merupakan sekumpulan method signature yang menentukan behavior yang dimiliki oleh interface.
// Interface dapat berlaku sebagai kontrak yang harus dimiliki oleh suatu objek agar teridentifikasi sebagai interface tersebut.
// Implementasi interface pada golang sifatnya adalah implicit, atau dikenal juga sebagai duck-typing.
// Singkatnya, jika suatu objek berbunyi "quack" maka objek tersebut adalah bebek.

// Untuk mendeklarasikan interface, kita dapat menggunakan keyword interface.
// Sebagai contoh, potongan kode berikut merupakan interface FoodSearcher yang memiliki dua method signature
//  FindByName(name string) (Food, error) error dan FindByOrigin(origin string) ([]Food, error)
type FoodSearcher interface {
	FindByName(name string) (Food, error)
	FindByOrigin(origin string) ([]Food, error)
}

type Food struct {
	Name   string
	Origin string
}

// Untuk mengimplementasikan interface tersebut, suatu objek harus memiliki kedua method signature yang dimiliki oleh FoodSearcher
// Pada contoh berikut terdapat dua objek yang akan mengimplementasikan interface FoodSearcher

// Objek pertama yaitu FoodSearcherWithMap. Objek ini menyimpan data Food di dalam map. Pencarian data Food dilakukan dengan mencari elemen pada map.
type FoodSearcherWithMap struct {
	data map[string]Food
}

func (fs FoodSearcherWithMap) FindByName(name string) (Food, error) {
	fmt.Println("FindByName from FoodSearcherWithMap")
	// implementations..
	return Food{}, nil
}

func (fs FoodSearcherWithMap) FindByOrigin(origin string) ([]Food, error) {
	fmt.Println("FindByOrigin from FoodSearcherWithMap")
	// implementations..
	return nil, nil
}

// Objek kedua yaitu FoodSearcherWithMysql. Objek ini menyimpan data Food di dalam database mysql. Pencarian data Food dilakukan dengan melakukan query database.
type FoodSearcherWithMysql struct {
}

func (fs FoodSearcherWithMysql) FindByName(name string) (Food, error) {
	fmt.Println("FindByName from FoodSearcherWithMysql")
	// implementations..
	return Food{}, nil
}

func (fs FoodSearcherWithMysql) FindByOrigin(origin string) ([]Food, error) {
	fmt.Println("FindByOrigin from FoodSearcherWithMysql")
	// implementations..
	return nil, nil
}

func main() {
	// Inisiasi FoodSearcher
	var FoodSearcher FoodSearcher

	// FoodSearcher menggunakan map
	FoodSearcher = FoodSearcherWithMap{
		data: map[string]Food{},
	}
	FoodSearcher.FindByName("Nasi Goreng")

	// FoodSearcher menggunakan MySQL
	FoodSearcher = FoodSearcherWithMysql{}
	FoodSearcher.FindByName("Nasi Goreng")
}
