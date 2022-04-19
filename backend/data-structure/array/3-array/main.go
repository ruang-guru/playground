// Diberikan array bilangan dengan variabel nums. Kita dapat mendefinisikan jumlah array berjalan sebagai runningSum[i] = sum(nums[0]…nums[i]).
// Kembalikan jumlah angka yang berjalan.
//
// Contoh 1
// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]
// Explanation: runningSum merupakan hasil dari [1, 1+2, 1+2+3, 1+2+3+4]
//
// Contoh 2
// Input: nums = [1,1,1,1,1]
// Output: [1,2,3,4,5]
// Explanation: runningSum merupakan hasil dari [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1]
//
// Contoh 3
// Input: nums = [3,1,2,10,1]
// Output: [3,4,6,16,17]

package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
