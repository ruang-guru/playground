package compareFast_test

import (
	"strconv"
	"testing"

	"github.com/ruang-guru/playground/backend/benchmarking-and-compiler/compare"
)

func BenchmarkUnmarshal(b *testing.B) {

	for n := 0; n < b.N; n++ {
		// use strconv.Itoa(n) to convert n to string, use UnmarshallSTD here
		// TODO: answer here
	}

}
