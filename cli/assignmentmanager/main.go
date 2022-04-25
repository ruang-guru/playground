package assignmentmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	graderapiclient "github.com/ruang-guru/playground/cli/assignmentmanager/client/graderapi"
	"github.com/ruang-guru/playground/cli/assignmentmanager/entity"
)

type assignmentSubmitter interface {
	GetAssignmentsByCourses([]string) ([]entity.Assignment, error)
	CreateAssignment(*entity.Assignment) (*entity.Assignment, error)
	UpdateAssignment(string, *entity.Assignment) error
}

func ManageAssignmentsFromJsonFile(jsonFilePath, graderEmail, graderPassword string) error {
	ctx := context.Background()
	graderApiClient := graderapiclient.NewGraderApiClientFromEmailAndPassword(ctx, graderEmail, graderPassword)
	tokenData, err := graderApiClient.Login()
	if err != nil {
		return fmt.Errorf("failed authenticating to server: %s", err.Error())
	}
	graderApiClient.SetTokenData(tokenData)

	localAssignments, err := LoadAssignmentsFromJsonFile(jsonFilePath)
	if err != nil {
		return fmt.Errorf("failed to load assignments from file %s: %s", jsonFilePath, err.Error())
	}

	remoteAssignments, err := GetRemoteAssignments(graderApiClient, GetCoursesFromAssignments(localAssignments))
	if err != nil {
		return fmt.Errorf("failed to get remote assignments: %s", err.Error())
	}

	assignmentsToBeCreated, assignmentsToBeUpdated := GetAssignmentsToBeCreatedAndUpdated(localAssignments, remoteAssignments)

	if len(assignmentsToBeCreated) == 0 && len(assignmentsToBeUpdated) == 0 {
		fmt.Println("no change in assignment, skipping tasks")
		return nil
	}

	if err := CreateNewAssignments(graderApiClient, assignmentsToBeCreated); err != nil {
		return fmt.Errorf("error %s", err.Error())
	}

	if err := UpdateAssignments(graderApiClient, assignmentsToBeUpdated); err != nil {
		return fmt.Errorf("error %s", err.Error())
	}

	return nil
}

func LoadAssignmentsFromJsonFile(jsonFilePath string) ([]entity.Assignment, error) {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return []entity.Assignment{}, fmt.Errorf("failed opening file %s: %s", jsonFilePath, err.Error())
	}
	defer jsonFile.Close()

	assignments, err := ParseAssignmentsFromJsonFile(jsonFile)
	if err != nil {
		return []entity.Assignment{}, fmt.Errorf("failed to read from file %s: %s", jsonFilePath, err.Error())
	}

	return assignments, nil
}

func ParseAssignmentsFromJsonFile(reader io.Reader) ([]entity.Assignment, error) {
	fileContentBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return []entity.Assignment{}, err
	}
	var assignments entity.Assignments
	json.Unmarshal(fileContentBytes, &assignments)

	return assignments.Assignments, nil
}

func GetCoursesFromAssignments(asgmts []entity.Assignment) []string {
	courseMap := make(map[string]bool)
	for _, assignment := range asgmts {
		if _, ok := courseMap[assignment.Course]; !ok {
			courseMap[assignment.Course] = true
		}
	}

	courses := make([]string, 0)
	for course := range courseMap {
		courses = append(courses, course)
	}
	return courses
}

func GetRemoteAssignments(graderApiClient assignmentSubmitter, courses []string) ([]entity.Assignment, error) {
	if len(courses) == 0 {
		return []entity.Assignment{}, nil
	}

	// TODO: change implementation when api GetAssignments is ready
	assignmentsOnServer, err := graderApiClient.GetAssignmentsByCourses(courses)
	if err != nil {
		return []entity.Assignment{}, err
	}
	return assignmentsOnServer, nil
}

func GetAssignmentsToBeCreatedAndUpdated(localAsgmts, remoteAsgmts []entity.Assignment) ([]*entity.Assignment, map[string]*entity.Assignment) {
	assignmentsToBeCreated := make([]*entity.Assignment, 0)
	assignmentsToBeUpdated := make(map[string]*entity.Assignment)

	remoteAsgmtsMap := make(map[string]*entity.Assignment)
	for idx, assignment := range remoteAsgmts {
		remoteAsgmtsMap[assignment.Path] = &remoteAsgmts[idx]
	}

	for idx, assignment := range localAsgmts {
		if _, exists := remoteAsgmtsMap[assignment.Path]; !exists {
			assignmentsToBeCreated = append(assignmentsToBeCreated, &localAsgmts[idx])
			continue
		}

		if !assignment.IsEquivalent(remoteAsgmtsMap[assignment.Path]) {
			assignmentsToBeUpdated[remoteAsgmtsMap[assignment.Path].Serial] = &localAsgmts[idx]
		}
	}

	return assignmentsToBeCreated, assignmentsToBeUpdated
}

func CreateNewAssignments(graderApiClient assignmentSubmitter, assignments []*entity.Assignment) error {
	for _, assignment := range assignments {
		fmt.Printf("creating new assignment: [path: %s; subject: %s; course: %s; weight: %f]\n", assignment.Path, assignment.Subject, assignment.Course, assignment.Weight)
		if _, err := graderApiClient.CreateAssignment(assignment); err != nil {
			return fmt.Errorf("failed to create assignment at path %s [%s]", assignment.Path, err.Error())
		}
	}
	return nil
}

func UpdateAssignments(graderApiClient assignmentSubmitter, assignmentMap map[string]*entity.Assignment) error {
	// TODO: improve by using go concurrency & errgroup ?
	for serial, assignment := range assignmentMap {
		fmt.Printf("updating assignment with serial %s: [path: %s; subject: %s; course: %s; weight: %f]\n", serial, assignment.Path, assignment.Subject, assignment.Course, assignment.Weight)
		if err := graderApiClient.UpdateAssignment(serial, assignment); err != nil {
			return fmt.Errorf("failed to update assignment %s [%s]", serial, err.Error())
		}
	}

	return nil
}
