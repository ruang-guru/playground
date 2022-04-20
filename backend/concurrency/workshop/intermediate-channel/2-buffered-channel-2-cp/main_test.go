package main

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
)

var _ = Describe("Buffered Channel", func() {
	When("buffer not full and receiver not ready", func() {
		It("won't fail", func() {
			result := make(chan string)
			experiment := gmeasure.NewExperiment("buffered-channel")
			AddReportEntry(experiment.Name, experiment)
			experiment.Sample(func(idx int) {
				experiment.MeasureDuration("buffered-channel", func() {
					go testBufferedChannel(result)
					for _, name := range names {
						Expect(<-result).To(Equal(fmt.Sprintf("%s say hello to %s", person, name)))
					}

				})
			}, gmeasure.SamplingConfig{N: 3})
			goroutineStats := experiment.GetStats("buffered-channel")
			meanDuration := goroutineStats.DurationFor(gmeasure.StatMean)
			Expect(meanDuration).To(BeNumerically("~", 100*time.Millisecond, 20*time.Millisecond))
		})
	})
})
