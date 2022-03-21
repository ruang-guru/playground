package answerremover_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAnswerremover(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Answerremover Suite")
}
