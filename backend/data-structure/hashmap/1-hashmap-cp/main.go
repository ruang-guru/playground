package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	ans := make(map[int]int)

	for _, val := range people {
		ans[val.age]++
	}
	return ans // TODO: replace this
}

func FilterByAge(people []Person, age int) []Person {
	ans := make([]Person, 0)

	for _, val := range people {
		if val.age == age {
			ans = append(ans, val)
		}
	}

	return ans // TODO: replace this
}
