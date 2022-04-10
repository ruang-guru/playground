// Diberikan string "jewels" yang mewakili jenis batu yang merupakan permata, dan string "stones" yang mewakili batu yang dimiliki.
// Setiap karakter dalam "stones" adalah jenis batu yang dimiliki.
// Kita ingin tahu berapa banyak batu yang dimiliki yang juga termasuk permata.
// Huruf peka (case sensitive) terhadap huruf besar-kecil, jadi "a" dianggap sebagai jenis batu yang berbeda dari "A".
//
// Contoh 1
// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3
// Explanation: ada 3 permata di dalam variabel stones, yaitu a, A, dan A
//
// Contoh 2
// Input: jewels = "z", stones = "ZZ"
// Output: 0
// Explanation: karena case sensitive, tidak ada permata di dalam variabel stones

package main

import "fmt"

func main() {
	var jewels = "aA"
	var stones = "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
}

func numJewelsInStones(jewels string, stones string) int {
	stoneMap := map[rune]bool{}
	count := 0

	for _, val := range jewels {
		stoneMap[val] = true
	}

	for _, stone := range stones {
		if _, exists := stoneMap[stone]; exists {
			count++
		}
	}

	return count
}
