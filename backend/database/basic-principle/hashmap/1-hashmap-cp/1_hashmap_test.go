package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/hashmap/1-hashmap-cp"
)

var _ = Describe("Hashmap", func() {
	It("can find single value", func() {
		hm := bdp.NewHashMap()
		hm.Put(1, "value1")

		val, _ := hm.Get(1)
		Expect(val).To(Equal("value1"))
	})

	Context("get ranges of keys equivalent to sql `SELECT * FROM table WHERE key >= left_range AND key <= right_range`", func() {
		Measure("It doesn't get range efficiently", func(b Benchmarker) {
			hm := bdp.NewHashMap()
			for i := 0; i < 100000; i++ {
				hm.Put(i+100000, fmt.Sprintf("i"))
			}
			var res []string
			runtime := b.Time("runtime", func() {
				res, _ = hm.GetRange(0, 100000)
			})
			Expect(runtime.Seconds()).ToNot(BeNumerically("<", 0.001))
			Expect(len(res)).To(Equal(1))
		}, 3)
	})
})
