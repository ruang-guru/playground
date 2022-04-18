package main

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Goroutine", func() {
	It("can finish task faster", func() {
		experiment := gmeasure.NewExperiment("chess")
		AddReportEntry(experiment.Name, experiment)
		experiment.Sample(func(idx int) {
			experiment.MeasureDuration("chess", func() {
				startTournament()
			})
		}, gmeasure.SamplingConfig{N: 3})
		goroutineStats := experiment.GetStats("chess")
		meanDuration := goroutineStats.DurationFor(gmeasure.StatMean)
		Expect(meanDuration).To(BeNumerically("~", 150*time.Millisecond, 50*time.Millisecond))
	})
})
