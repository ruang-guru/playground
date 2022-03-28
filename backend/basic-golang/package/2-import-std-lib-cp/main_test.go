package main

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountDays", func() {
	It("returns 6 when start date is 21 March 2022 and end date is 27 March 2022", func() {
		start := time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC)
		end := time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC)
		dayDifference := CountDays(start, end)
		Expect(dayDifference).To(Equal(6))
	})
	It("returns -1 when start date is 21 March 2022 and end date is 20 March 2022", func() {
		start := time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC)
		end := time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC)
		dayDifference := CountDays(start, end)
		Expect(dayDifference).To(Equal(-1))
	})
	It("returns 0 when start date and end date is same day", func() {
		start := time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC)
		end := time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC)
		dayDifference := CountDays(start, end)
		Expect(dayDifference).To(Equal(0))
	})
})
