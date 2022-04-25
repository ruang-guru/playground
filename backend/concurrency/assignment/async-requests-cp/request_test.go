package requests_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	req "github.com/ruang-guru/playground/backend/concurrency/assignment/async-requests-cp/client/request"
	"github.com/ruang-guru/playground/backend/concurrency/assignment/async-requests-cp/server/handler"
)

var _ = Describe("Request", func() {
	When("running sequentially", func() {
		It("less efficient", func() {
			ts := httptest.NewServer(http.HandlerFunc(handler.GetMessage))
			urls := []string{}
			defer ts.Close()
			for i := 0; i < 10; i++ {
				urls = append(urls, ts.URL)
			}
			experiment := gmeasure.NewExperiment("async-request")
			AddReportEntry(experiment.Name, experiment)
			experiment.Sample(func(idx int) {
				experiment.MeasureDuration("async-request", func() {
					req.SyncHttpGets(urls)

				})
			}, gmeasure.SamplingConfig{N: 3})
			goroutineStats := experiment.GetStats("async-request")
			meanDuration := goroutineStats.DurationFor(gmeasure.StatMean)
			Expect(meanDuration).To(BeNumerically(">", 20*time.Millisecond))
		})
	})

	When("running concurrently", func() {
		It("more efficient", func() {
			ts := httptest.NewServer(http.HandlerFunc(handler.GetMessage))
			defer ts.Close()
			urls := []string{}
			for i := 0; i < 10; i++ {
				urls = append(urls, ts.URL)
			}
			experiment := gmeasure.NewExperiment("async-request")
			AddReportEntry(experiment.Name, experiment)
			experiment.Sample(func(idx int) {
				experiment.MeasureDuration("async-request", func() {
					req.AsyncHttpGets(urls)
				})
			}, gmeasure.SamplingConfig{N: 3})
			goroutineStats := experiment.GetStats("async-request")
			meanDuration := goroutineStats.DurationFor(gmeasure.StatMean)
			Expect(meanDuration).To(BeNumerically("<", 20*time.Millisecond))
		})
	})

})
