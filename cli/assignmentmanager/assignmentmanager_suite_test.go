package assignmentmanager_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAssignmentmanager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assignmentmanager Suite")
}
