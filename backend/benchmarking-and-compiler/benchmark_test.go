package main

import (
	"os/exec"
	"runtime"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = AfterSuite(func() {
	// clean fast.out and std.out
	os := runtime.GOOS
	var err error
	if os == "windows" {
		err = exec.Command("cmd", "/C", "del", "fast.txt", "std.txt").Run()
	} else {
		err = exec.Command("rm", "fast.txt", "std.txt").Run()
	}
	Expect(err).To(BeNil())
})

var _ = Describe("Benchmark", func() {
	When("you implement BenchmarkUnmarshal in std-benchmark and BenchmarkUnmarshalFast in fast-benchmark", func() {
		It("should be able to generate report benchmark report file std.out and fast.out", func() {
			By("running command to generate the file")
			err := RunBenchmark("./std-benchmark", "std.txt")
			Expect(err).To(BeNil())
			fileName := FileExist("std.txt")
			Expect(fileName).To(Equal("std.txt"))

			err = RunBenchmark("./fast-benchmark", "fast.txt")
			Expect(err).To(BeNil())
			fileName = FileExist("fast.txt")
			Expect(fileName).To(Equal("fast.txt"))
		})
	})

	When("you have std.out and fast.out file we can benchcmp them", func() {
		It("should generate a .txt file that have report on it", func() {
			err := RunBenchcmp()
			Expect(err).To(BeNil())
			fileName := FileExist("report.txt")
			Expect(fileName).To(Equal("report.txt"))
		})
	})

	Describe("report.txt", func() {
		It("what you can get by reading", func() {
			answer := `pada report tersebut kita dapat mengetahui bahwa library fastJson 37% lebih cepat dari std library dan mempunyai 19.8
			lebih sedikit allocation tetapi memakan sekitar 123%`
			Expect(answer).To(Not(BeEmpty()))
		})
	})
})
