// Tulis program sebagai fungsi yang mengembalikan union dari dua set.
// Jika kedua set union, fungsi tersebut harus mengembalikan hasilnya.
//
// Contoh 1
// Input: a = [1, 2, 3, 4, 5], b = [2, 3, 5, 7, 11]
// Output: [1, 2, 3, 4, 5, 7, 11]
// Explanation: intersection dari a dan b adalah 2, 3 dan 5
//
// Contoh 2
// Input: a = [1, 2, 3], b = [4, 5, 6]
// Output: [1, 2, 3, 4, 5, 6]
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var int1 = []int{1, 2, 3, 4, 5}
	var int2 = []int{2, 3, 5, 7, 11}
	fmt.Println(union(int1, int2))
}

func union(int1, int2 []int) []int {
	hash := make(map[int]bool)

	for _, item := range int1 {
		hash[item] = true
	}

	for _, item := range int2 {
		if _, exist := hash[item]; !exist {
			int1 = append(int1, item)
		}
	}
	return int1
}
