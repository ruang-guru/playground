package graderapiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/ruang-guru/playground/cli/assignmentmanager/entity"
)

type CreateAssignmentResponse struct {
	Data    entity.Assignment `json:"data"`
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Detail  string            `json:"detail"`
}

type AssignmentListResponse struct {
	Assignment []entity.Assignment `json:"assignments"`
}

type GetAssignmentsResponse struct {
	Data    AssignmentListResponse `json:"data"`
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Detail  string                 `json:"detail"`
}

func (g *GraderApiClient) CreateAssignment(assignment *entity.Assignment) (*entity.Assignment, error) {
	totalTestCase, err := strconv.Atoi(assignment.TotalTestCase)
	if err != nil {
		return nil, err
	}

	requestData := map[string]interface{}{
		"path":          assignment.Path,
		"subject":       assignment.Subject,
		"course":        assignment.Course,
		"weight":        assignment.Weight,
		"totalTestCase": totalTestCase,
	}

	postBody, _ := json.Marshal(requestData)

	reqBody := bytes.NewBuffer(postBody)

	resp, err := g.AuthenticatedPost(GraderApiUrl+"/assignment", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var createAssignmentResponse CreateAssignmentResponse
	if err := json.NewDecoder(resp.Body).Decode(&createAssignmentResponse); err != nil {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to parse response body\n%s\n\n%s", string(responseBody), err.Error())
	}

	if createAssignmentResponse.Status != "success" {
		return nil, fmt.Errorf("[%s] %s", createAssignmentResponse.Message, createAssignmentResponse.Detail)
	}

	return &createAssignmentResponse.Data, nil
}

// TODO: implement
func (g *GraderApiClient) UpdateAssignment(serial string, assignment *entity.Assignment) error {
	totalTestCase := 0
	var err error

	if assignment.TotalTestCase != "" {
		totalTestCase, err = strconv.Atoi(assignment.TotalTestCase)
		if err != nil {
			return err
		}
	}

	requestData := map[string]interface{}{
		"subject":       assignment.Subject,
		"course":        assignment.Course,
		"weight":        assignment.Weight,
		"totalTestCase": totalTestCase,
	}

	putBody, _ := json.Marshal(requestData)

	reqBody := bytes.NewBuffer(putBody)

	resp, err := g.AuthenticatedPut(GraderApiUrl+"/assignment/"+serial, reqBody)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var createAssignmentResponse CreateAssignmentResponse
	if err := json.NewDecoder(resp.Body).Decode(&createAssignmentResponse); err != nil {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to parse response body\n%s\n\n%s", string(responseBody), err.Error())
	}

	if createAssignmentResponse.Status != "success" {
		return fmt.Errorf("[%s] %s", createAssignmentResponse.Message, createAssignmentResponse.Detail)
	}

	return nil
}

func (g *GraderApiClient) GetAssignmentsByCourse(course string) ([]entity.Assignment, error) {
	resp, err := g.AuthenticatedGet(fmt.Sprintf("%s/assignment/%s", GraderApiUrl, course))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var getAssignmentsResponse GetAssignmentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&getAssignmentsResponse); err != nil {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to parse response body\n%s\n\n%s", string(responseBody), err.Error())
	}

	if getAssignmentsResponse.Status != "success" {
		return nil, fmt.Errorf("[%s] %s", getAssignmentsResponse.Message, getAssignmentsResponse.Detail)
	}

	return getAssignmentsResponse.Data.Assignment, nil
}

func (g *GraderApiClient) GetAssignmentsByCourses(courses []string) ([]entity.Assignment, error) {
	assignments := make([]entity.Assignment, 0)
	for _, course := range courses {
		additionalAssignments, err := g.GetAssignmentsByCourse(course)
		if err != nil {
			return []entity.Assignment{}, err
		}
		assignments = append(assignments, additionalAssignments...)
	}

	return assignments, nil
}
