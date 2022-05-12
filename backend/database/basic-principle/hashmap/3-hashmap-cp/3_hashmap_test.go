package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/hashmap/3-hashmap-cp"
)

var _ = Describe("Array", func() {
	Describe("get data related by Name equivalent to sql `SELECT * FROM table WHERE name = params_name`", func() {
		It("should get row table related by Name", func() {
			hm := bdp.NewUser()
			hm.Insert("Gina", 18)

			w := hm.WhereByName(bdp.SecondaryKey("Gina"))
			Expect(w).To(HaveLen(1))
			Expect(w[0].Name).To(Equal(bdp.SecondaryKey("Gina")))
			Expect(w[0].Age).To(Equal(18))

			hm.Insert("Juno", 19)
			w = hm.WhereByName(bdp.SecondaryKey("Juno"))
			Expect(w).To(HaveLen(1))
			Expect(w[0].Name).To(Equal(bdp.SecondaryKey("Juno")))
			Expect(w[0].Age).To(Equal(19))
		})
	})

	Describe("get data begin with letter equivalent to sql `SELECT * FROM table WHERE name LIKE '<letter_param>%';`", func() {
		It("should get row table related by start with letter_param", func() {
			hm := bdp.NewUser()
			hm.Insert("Gina", 18)
			hm.Insert("Juno", 19)
			hm.Insert("Gina", 20)
			hm.Insert("Juno", 21)
			hm.Insert("Gina", 22)
			hm.Insert("Juno", 23)
			hm.Insert("Gina", 24)

			w := hm.WhereNameBeginsWith("G")
			Expect(w).To(HaveLen(4))
			Expect(w[0].Name).To(Equal(bdp.SecondaryKey("Gina")))
			Expect(w[0].Age).To(Equal(18))
			Expect(w[1].Name).To(Equal(bdp.SecondaryKey("Gina")))
			Expect(w[1].Age).To(Equal(20))
			Expect(w[2].Name).To(Equal(bdp.SecondaryKey("Gina")))
			Expect(w[2].Age).To(Equal(22))
			Expect(w[3].Name).To(Equal(bdp.SecondaryKey("Gina")))
			Expect(w[3].Age).To(Equal(24))

			w = hm.WhereNameBeginsWith("J")
			Expect(w).To(HaveLen(3))
			Expect(w[0].Name).To(Equal(bdp.SecondaryKey("Juno")))
			Expect(w[0].Age).To(Equal(19))
			Expect(w[1].Name).To(Equal(bdp.SecondaryKey("Juno")))
			Expect(w[1].Age).To(Equal(21))
			Expect(w[2].Name).To(Equal(bdp.SecondaryKey("Juno")))
			Expect(w[2].Age).To(Equal(23))

		})
	})
})
