package main

import "fmt"

func main() {
	numList := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := int64(9)
	fmt.Println(BinarySearch(numList, key))
}

//Recursive Binary Search
func BinarySearch(numList []int64, key int64) int {
	low := 0
	high := len(numList) - 1

	if low <= high {
		// TODO: answer here
		mid := high / 2
		if key == numList[mid] {
			return 1
		}
		if key > numList[mid] {
			return BinarySearch(numList[mid+1:high+1], key)
		} else if key < numList[mid] {
			return BinarySearch(numList[0:mid], key)
		}
	}
	return 0
}
