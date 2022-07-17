package assignmentmanager_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ruang-guru/playground/cli/assignmentmanager"
	"github.com/ruang-guru/playground/cli/assignmentmanager/entity"
)

type dummyGraderClient struct {
	CountGetAssignmentsByCoursesCall int
	CountCreateAssignmentCall        int
	CountUpdateAssignment            int

	CreatedAssignmentPaths   []string
	UpdatedAssignmentSerials map[string]bool
}

func NewDummyGraderClient() *dummyGraderClient {
	return &dummyGraderClient{
		CountGetAssignmentsByCoursesCall: 0,
		CountCreateAssignmentCall:        0,
		CountUpdateAssignment:            0,
		CreatedAssignmentPaths:           make([]string, 0),
		UpdatedAssignmentSerials:         make(map[string]bool),
	}
}

func (d *dummyGraderClient) GetAssignmentsByCourses([]string) ([]entity.Assignment, error) {
	d.CountGetAssignmentsByCoursesCall++
	return []entity.Assignment{}, nil
}

func (d *dummyGraderClient) CreateAssignment(a *entity.Assignment) (*entity.Assignment, error) {
	d.CountCreateAssignmentCall++
	d.CreatedAssignmentPaths = append(d.CreatedAssignmentPaths, a.Path)
	return &entity.Assignment{}, nil
}

func (d *dummyGraderClient) UpdateAssignment(s string, a *entity.Assignment) error {
	d.CountUpdateAssignment++
	d.UpdatedAssignmentSerials[s] = true
	return nil
}

var _ = Describe("Assignmenmanager", func() {
	Describe("parseAssignmentsFromJsonFile", func() {
		When("given correct json file", func() {
			It("should return correct list of assignments", func() {
				var buffer bytes.Buffer
				buffer.WriteString(`
				{
					"assignments": [
						{
							"path": "deny/test-01",
							"subject": "fe",
							"course": "frontend",
							"weight": 9.1,
							"deadline": "2022-04-08T15:51:00+07:00"
						},
						{
							"path": "deny/test-02",
							"subject": "fe",
							"course": "cypress",
							"deadline": "2022-04-08T15:51:00+07:00"
						},
						{
							"path": "deny/test-03",
							"subject": "be",
							"course": "backend",
							"deadline": "2022-04-08T15:51:00+07:00"
						},
						{
							"path": "deny/test-05",
							"subject": "be",
							"course": "frontend",
							"weight": 9.1,
							"deadline": "2022-04-08T15:51:00+07:00"
						}
					]
				}
				`)
				content, err := assignmentmanager.ParseAssignmentsFromJsonFile(&buffer)
				Expect(err).To(BeNil())
				Expect(len(content)).To(Equal(4))
			})
		})
	})

	Describe("getCoursesFromAssignments", func() {
		When("given slice of assignments", func() {
			It("should return correct set of assignments' courses", func() {
				assignments := []entity.Assignment{
					{
						Course: "course1",
					},
					{
						Course: "course1",
					},
					{
						Course: "course2",
					},
					{
						Course: "course3",
					},
				}

				courses := assignmentmanager.GetCoursesFromAssignments(assignments)

				Expect(len(courses)).To(Equal(3))
			})
		})
		When("given empty slice", func() {
			It("should return empty slice of string", func() {
				courses := assignmentmanager.GetCoursesFromAssignments([]entity.Assignment{})
				Expect(courses).To(Equal([]string{}))
			})
		})
	})

	Describe("getRemoteAssignments", func() {
		When("called with non empty slice of course", func() {
			It("should invoke API call to GetAssignmentsByCourses", func() {
				d := NewDummyGraderClient()
				assignmentmanager.GetRemoteAssignments(d, []string{"backend", "frontend"})
				assignmentmanager.GetRemoteAssignments(d, []string{"backend", "frontend"})
				assignmentmanager.GetRemoteAssignments(d, []string{"backend", "frontend"})

				Expect(d.CountGetAssignmentsByCoursesCall).To(Equal(3))
			})
		})

		When("called with empty slice of course", func() {
			It("should not invoke any API call", func() {
				d := NewDummyGraderClient()
				assignmentmanager.GetRemoteAssignments(d, []string{})

				Expect(d.CountGetAssignmentsByCoursesCall).To(Equal(0))
			})
		})
	})

	Describe("getAssignmentsToBeCreatedAndUpdated", func() {
		When("there are new assignments", func() {
			It("should return correct list of new assignments on the first return value", func() {
				loc := []entity.Assignment{
					{Path: "asgmt2"},
					{Path: "asgmt3"},
				}
				rem := []entity.Assignment{
					{Path: "asgmt1"},
				}

				newAssignments, _ := assignmentmanager.GetAssignmentsToBeCreatedAndUpdated(loc, rem)

				Expect(len(newAssignments)).To(Equal(2))
				Expect(newAssignments[0].Path).To(Equal("asgmt2"))
				Expect(newAssignments[1].Path).To(Equal("asgmt3"))
			})
		})

		When("there is no new assignment", func() {
			It("should return empty slice on the first return value", func() {
				loc := []entity.Assignment{
					{Path: "asgmt1"},
				}
				rem := []entity.Assignment{
					{Path: "asgmt1"},
					{Path: "asgmt2"},
					{Path: "asgmt3"},
				}

				newAssignments, _ := assignmentmanager.GetAssignmentsToBeCreatedAndUpdated(loc, rem)

				Expect(len(newAssignments)).To(Equal(0))
			})
		})

		When("there are assignments to be updated", func() {
			It("should return correct map of assignments on the second return value", func() {
				loc := []entity.Assignment{
					{Path: "asgmt1", Weight: 10},
					{Path: "asgmt2", Weight: 20},
					{Path: "asgmt3", Weight: 30},
				}
				rem := []entity.Assignment{
					{Serial: "asgmt1", Path: "asgmt1", Weight: 0},
					{Serial: "asgmt2", Path: "asgmt2", Weight: 10},
					{Serial: "asgmt3", Path: "asgmt3"},
					{Serial: "asgmt4", Path: "asgmt4"},
				}

				_, assignmentToBeUpdated := assignmentmanager.GetAssignmentsToBeCreatedAndUpdated(loc, rem)
				Expect(len(assignmentToBeUpdated)).To(Equal(3))
				Expect(assignmentToBeUpdated["asgmt1"].Weight).To(Equal(float32(10)))
				Expect(assignmentToBeUpdated["asgmt2"].Weight).To(Equal(float32(20)))
				Expect(assignmentToBeUpdated["asgmt3"].Weight).To(Equal(float32(30)))
			})
		})

		When("there is no assignment to be updated", func() {
			It("should return empty map", func() {
				loc := []entity.Assignment{
					{Path: "asgmt1", Weight: 10},
					{Path: "asgmt2", Weight: 20},
					{Path: "asgmt3", Weight: 30},
				}
				rem := []entity.Assignment{
					{Serial: "asgmt1", Path: "asgmt1", Weight: 10},
					{Serial: "asgmt2", Path: "asgmt2", Weight: 20},
					{Serial: "asgmt4", Path: "asgmt4"},
				}

				_, assignmentToBeUpdated := assignmentmanager.GetAssignmentsToBeCreatedAndUpdated(loc, rem)

				Expect(len(assignmentToBeUpdated)).To(Equal(0))
			})
		})

		When("there are both new assignments and assignments to be updated", func() {
			It("should return correct map of assignments on the second return value", func() {
				loc := []entity.Assignment{
					{Path: "asgmt1", Weight: 10},
					{Path: "asgmt2", Weight: 20},
					{Path: "asgmt3", Weight: 30},
					{Path: "asgmt5", Weight: 30},
					{Path: "asgmt6", Weight: 70},
				}
				rem := []entity.Assignment{
					{Serial: "asgmt1", Path: "asgmt1", Weight: 0},
					{Serial: "asgmt2", Path: "asgmt2", Weight: 10},
					{Serial: "asgmt3", Path: "asgmt3"},
					{Serial: "asgmt4", Path: "asgmt4"},
				}

				newAssignments, assignmentToBeUpdated := assignmentmanager.GetAssignmentsToBeCreatedAndUpdated(loc, rem)

				Expect(len(newAssignments)).To(Equal(2))
				Expect(newAssignments[0].Path).To(Equal("asgmt5"))
				Expect(newAssignments[1].Path).To(Equal("asgmt6"))

				Expect(len(assignmentToBeUpdated)).To(Equal(3))
				Expect(assignmentToBeUpdated["asgmt1"].Weight).To(Equal(float32(10)))
				Expect(assignmentToBeUpdated["asgmt2"].Weight).To(Equal(float32(20)))
				Expect(assignmentToBeUpdated["asgmt3"].Weight).To(Equal(float32(30)))
			})
		})
	})

	Describe("CreateNewAssignments", func() {
		When("called with slice of assignments", func() {
			It("should create assignments with correct order", func() {
				d := NewDummyGraderClient()

				assignmentmanager.CreateNewAssignments(d, []*entity.Assignment{
					{Path: "asgmt1"},
					{Path: "asgmt2"},
					{Path: "asgmt3"},
					{Path: "asgmt4"},
				})

				Expect(d.CountCreateAssignmentCall).To(Equal(4))
				Expect(d.CreatedAssignmentPaths[0]).To(Equal("asgmt1"))
				Expect(d.CreatedAssignmentPaths[1]).To(Equal("asgmt2"))
				Expect(d.CreatedAssignmentPaths[2]).To(Equal("asgmt3"))
				Expect(d.CreatedAssignmentPaths[3]).To(Equal("asgmt4"))
			})
		})
	})

	Describe("UpdateAssignments", func() {
		When("called with slice of assignments", func() {
			It("should update all the assignments specified", func() {
				d := NewDummyGraderClient()

				assignmentmanager.UpdateAssignments(d, map[string]*entity.Assignment{
					"asgmt1": {Path: "asgmt1"},
					"asgmt2": {Path: "asgmt2"},
					"asgmt3": {Path: "asgmt3"},
				})

				Expect(d.CountUpdateAssignment).To(Equal(3))
				Expect(d.UpdatedAssignmentSerials["asgmt1"]).To(BeTrue())
				Expect(d.UpdatedAssignmentSerials["asgmt2"]).To(BeTrue())
				Expect(d.UpdatedAssignmentSerials["asgmt3"]).To(BeTrue())
			})
		})
	})
})
