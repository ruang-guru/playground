package main

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Goroutine", func() {
	It("can make program run concurrently", func() {
		experiment := gmeasure.NewExperiment("goroutine")
		AddReportEntry(experiment.Name, experiment)
		calledAndy := false
		calledBob := false
		experiment.Sample(func(idx int) {
			experiment.MeasureDuration("goroutine", func() {
				call(&calledAndy, &calledBob)
			})
		}, gmeasure.SamplingConfig{N: 3})
		goroutineStats := experiment.GetStats("goroutine")
		meanDuration := goroutineStats.DurationFor(gmeasure.StatMean)
		Expect(calledAndy && calledBob).To(Equal(true))
		Expect(meanDuration).To(BeNumerically("~", 200*time.Millisecond, 50*time.Millisecond))
	})
})
