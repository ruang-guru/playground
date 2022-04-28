package main

import (
	"math/rand"
	"testing"
)

func randArray(n int, maxVal int) [][]int {
	a := make([][]int, n)
	for i := range a {
		row := make([]int, n)
		a[i] = row
		for j := range row {
			row[j] = rand.Intn(maxVal + 1)
		}
	}
	return a
}

func sumRowOrder(a [][]int) int {
	total := 0
	for _, row := range a {
		for _, x := range row {
			total += x
		}
	}
	return total
}

func sumColumnOrder(a [][]int) int {
	total := 0
	size := len(a)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			total += a[y][x]
		}
	}
	return total
}

func BenchmarkSumColumn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumColumnOrder(randArray(1000, 10))
	}
}

func BenchmarkSumRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumRowOrder(randArray(1000, 10))
	}
}
