package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/design-patterns/creational/1-factory-pattern-cp/factoryLanjutan"
)

var _ = Describe("Factory", func() {
	Describe("Factory with schedule", func() {
		Context("Netflix", func() {
			It("Should returned content with respected value based on schedule", func() {
				var contentCreator factoryLanjutan.ContentCreator
				contentCreator = &factoryLanjutan.NetflixKorea{}

				content := contentCreator.Produce(factoryLanjutan.Monday)
				contentName := content.Play()
				Expect(contentName).To(Equal("Our Beloved Summer"))

				content = contentCreator.Produce(factoryLanjutan.Sunday)
				contentName = content.Play()
				Expect(contentName).To(Equal("All of Us Are Dead"))

				content = contentCreator.Produce(factoryLanjutan.Tuesday)
				contentName = content.Play()
				Expect(contentName).To(Equal("Start Up"))
			})

		})

	})
})
