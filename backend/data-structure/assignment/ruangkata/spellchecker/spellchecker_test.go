package spellchecker_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/ruangkata/spellchecker"
)

var _ = Describe("Spellchecker", func() {
	Describe("CheckWord", func() {
		It("correctly check valid word", func() {
			sc, err := spellchecker.NewEnglishSpellChecker()
			Expect(err).NotTo(HaveOccurred())

			ok := sc.CheckWord("hello")
			Expect(ok).To(BeTrue())

			ok = sc.CheckWord("world")
			Expect(ok).To(BeTrue())
		})
		It("correctly check invalid word", func() {
			sc, err := spellchecker.NewEnglishSpellChecker()
			Expect(err).NotTo(HaveOccurred())

			ok := sc.CheckWord("halow")
			Expect(ok).To(BeFalse())

			ok = sc.CheckWord("dunia")
			Expect(ok).To(BeFalse())
		})

		It("is case insensitive", func() {
			sc, err := spellchecker.NewEnglishSpellChecker()
			Expect(err).NotTo(HaveOccurred())

			ok := sc.CheckWord("HellO")
			Expect(ok).To(BeTrue())

			ok = sc.CheckWord("america")
		})
	})

	Describe("CheckSentences", func() {
		It("split valid and invalid words", func() {
			sc, err := spellchecker.NewEnglishSpellChecker()
			Expect(err).NotTo(HaveOccurred())

			valid, invalid := sc.CheckSentence("hello world")
			Expect(valid).To(Equal([]string{"hello", "world"}))
			Expect(invalid).To(Equal([]string{}))

			valid, invalid = sc.CheckSentence("hello world this is a sentence")
			Expect(valid).To(Equal([]string{"hello", "world", "this", "is", "a", "sentence"}))
			Expect(invalid).To(Equal([]string{}))

			valid, invalid = sc.CheckSentence("hello world this is a sentence tapi mengandung kata2 nggak bener")
			Expect(valid).To(Equal([]string{"hello", "world", "this", "is", "a", "sentence"}))
			Expect(invalid).To(Equal([]string{"tapi", "mengandung", "kata2", "nggak", "bener"}))
		})
	})
})
