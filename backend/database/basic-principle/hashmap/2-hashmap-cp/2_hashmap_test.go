package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bdp "github.com/ruang-guru/playground/backend/database/basic-principle/hashmap/2-hashmap-cp"
)

var _ = Describe("Array", func() {
	Describe("get all data related by ID equivalent to sql `SELECT * FROM table WHERE id = params_id`", func() {
		It("should get row table related by ID", func() {
			hm := bdp.NewInvoice()
			hm.Insert("PA-100", "Gina", "Surabaya", "081212121212")
			where := hm.Where(1)

			Expect(where.SubscriptionCode).To(Equal("PA-100"))
			Expect(where.CustomerName).To(Equal("Gina"))
			Expect(where.Address).To(Equal("Surabaya"))
			Expect(where.Phone).To(Equal("081212121212"))

			hm.Insert("PA-200", "Doni", "Jakarta", "084324234234")
			where = hm.Where(2)
			Expect(where.SubscriptionCode).To(Equal("PA-200"))
			Expect(where.CustomerName).To(Equal("Doni"))
			Expect(where.Address).To(Equal("Jakarta"))
			Expect(where.Phone).To(Equal("084324234234"))
		})
	})

	Describe("update data related by ID equivalent to sql `UPDATE table SET SubscriptionCode= 'PA-200', CustomerName= 'Doni', Address= 'Jakarta', Phone= '084324234234'  WHERE ID = 1;`", func() {
		It("should updated row table related by ID", func() {
			hm := bdp.NewInvoice()
			hm.Insert("PA-100", "Gina", "Surabaya", "081212121212")

			res, err := hm.Update(1, "PA-200", "Doni", "Jakarta", "084324234234")
			Expect(err).To(BeNil())
			Expect(res.SubscriptionCode).To(Equal("PA-200"))
			Expect(res.CustomerName).To(Equal("Doni"))
			Expect(res.Address).To(Equal("Jakarta"))
			Expect(res.Phone).To(Equal("084324234234"))
		})
	})
})
