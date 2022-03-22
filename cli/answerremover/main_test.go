package answerremover_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ruang-guru/playground/cli/answerremover"
)

var _ = Describe("Answerremover", func() {
	Describe("RemoveAnswersBlock", func() {
		When("there is no answer block indicator", func() {
			It("returns the same string", func() {
				s := "Hello, World!"

				Expect(answerremover.RemoveAnswerBlock(s)).To(Equal("Hello, World!"))
			})
		})

		It("removes text in the answer block", func() {
			s := `//beginanswer
Hello, World!
//endanswer`

			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`// TODO: answer here`))
		})

		When("answer block is not closed", func() {
			It("returns error", func() {
				s := `//beginanswer
Hello, World!`

				_, err := answerremover.RemoveAnswerBlock(s)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("answer block is not closed"))
			})
		})

		When("answer block is closed without opening pair", func() {
			It("returns error", func() {
				s := `//endanswer
Hello, World!`

				_, err := answerremover.RemoveAnswerBlock(s)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("found endanswer but no matched beginanswer"))
			})
		})

		When("answer block is opened multiple times", func() {
			It("returns error", func() {
				s := `//beginanswer
Hello, World!
//beginanswer
Hello, World!`

				_, err := answerremover.RemoveAnswerBlock(s)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("beginanswer found twice"))
			})
		})

		Describe("replacement with panic", func() {
			When("endanswer is appended with 'panic'", func() {
				It("replaces the answer block with panic code", func() {
					s := `//beginanswer
Hello, World!
//endanswer panic`

					replacement, _ := answerremover.RemoveAnswerBlock(s)
					Expect(replacement).To(Equal(`panic("Not yet implemented") // TODO: answer here`))
				})
				It("reserves the indentation", func() {
					s := `
	//beginanswer
	Hello, World!
	//endanswer panic`

					replacement, _ := answerremover.RemoveAnswerBlock(s)
					Expect(replacement).To(Equal("\n\tpanic(\"Not yet implemented\") // TODO: answer here"))
				})
			})
		})

		Describe("replacement with error", func() {
			When("endanswer is appended with 'error'", func() {
				It("replaces the answer block with error code", func() {
					s := `//beginanswer
Hello, World!
//endanswer error`

					replacement, _ := answerremover.RemoveAnswerBlock(s)
					Expect(replacement).To(Equal(`throw new Error("Method not implemented."); // TODO: answer here`))
				})
			})
		})

		Describe("replacement with code", func() {
			When("endanswer is appended with strings not in keyword", func() {
				It("replaces the answer block with the string", func() {
					s := `//beginanswer
Hello, World!
//endanswer fmt.Sprintf("%s", "Hello, World!")`

					replacement, _ := answerremover.RemoveAnswerBlock(s)
					Expect(replacement).To(Equal(`fmt.Sprintf("%s", "Hello, World!") // TODO: replace this`))
				})
			})
		})

		Describe("replacement with todo only", func() {
			It("adds #TODO and reserve the indentation", func() {
				s := `
	//beginanswer
	Hello, World!
	//endanswer`
				replacement, _ := answerremover.RemoveAnswerBlock(s)
				Expect(replacement).To(Equal("\n\t// TODO: answer here"))
			})
		})

		Describe("replacement with nothing", func() {
			It("just removes code & doesn't add #TODO", func() {
				s := `a := 2
	//beginanswer
	Hello, World!
	//endanswer nop`
				replacement, _ := answerremover.RemoveAnswerBlock(s)
				Expect(replacement).To(Equal("a := 2"))
			})
		})

		// 		Describe("replacement with no types (Typescript specific)", func() {
		// 			It("appends //TODO only in the line where the type is removed", func() {
		// 				s := `
		// //beginanswer
		// let a: number = 2
		// let b: string = "Hello, World!"
		// if (a === 2) {
		// }
		// //endanswer removetype`
		// 				replacement, _ := answerremover.RemoveAnswerBlock(s)
		// 				Expect(replacement).To(Equal(`
		// let a = 2 // TODO: add type here
		// let b = "Hello, World!" // TODO: add type here
		// if (a === 2) {
		// }`))
		// 			})
		// 		})

		It("accepts different variations of comment", func() {
			s := `//beginanswer
Hello, World!
//endanswer`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`// TODO: answer here`))

			s = `#beginanswer 
Hello, World!
#endanswer`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`# TODO: answer here`))

			s = `--beginanswer 
Hello, World!
--endanswer`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`-- TODO: answer here`))

			s = `/*beginanswer*/
Hello, World!
/*endanswer*/`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`/* TODO: answer here */`))

			s = `{/* beginanswer */}
Hello, World!
{/* endanswer */}`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`{/* TODO: answer here */}`))

			s = `<!-- beginanswer -->
Hello, World!
<!-- endanswer -->`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`<!-- TODO: answer here -->`))
		})

		It("tolerate spaces", func() {
			s := `// beginanswer
Hello, World!
// endanswer`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`// TODO: answer here`))

			s = `// beginanswer 
Hello, World!
// endanswer `
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`// TODO: answer here`))

			s = `//    beginanswer     
Hello, World!
//    endanswer    `
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`// TODO: answer here`))

			s = `/*    beginanswer  */  
Hello, World!
/*    endanswer     */    `
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`/* TODO: answer here */`))

			s = `{    /*    beginanswer  */   }
Hello, World!
{    /*    endanswer     */   }`
			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`{/* TODO: answer here */}`))

			s = `<!--    beginanswer  -->
Hello, World!
<!--    endanswer     -->`

			Expect(answerremover.RemoveAnswerBlock(s)).To(Equal(`<!-- TODO: answer here -->`))
		})
	})
})
