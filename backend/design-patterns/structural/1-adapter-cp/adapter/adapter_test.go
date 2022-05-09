package adapter_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/backend/design-patterns/structural/1-adapter-cp/adapter"
)

var _ = Describe("Adapter", func() {
	Context("when adapting a walkman to a kaset", func() {
		It("should play the kaset", func() {
			mp3 := adapter.Mp3{Data: []byte(`this is taylor swift song`)}
			walkman := adapter.Walkman{}

			ad := adapter.Mp3ToKasetAdapter{Adaptee: walkman}
			Expect(ad.Play(mp3)).To(Equal("this is taylor swift song"))
		})
	})
})
