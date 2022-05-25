package entity

import (
	"errors"
	"path"
	"strings"
)

type Assignments struct {
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Serial        string  `json:"serial"`
	Path          string  `json:"path"`
	Subject       string  `json:"subject"`
	Course        string  `json:"course"`
	Weight        float32 `json:"weight"`
	TotalTestCase string  `json:"totalTestCase"`
}

var (
	ErrInvalidPath = errors.New("invalid path, assignment path should be inside repository")
)

func NewAssignment(assignmentPath, subject, course string, weight float32, deadline string) (*Assignment, error) {
	cleanPath := path.Clean(assignmentPath)
	if strings.HasPrefix(cleanPath, ".") {
		return nil, ErrInvalidPath
	}
	cleanPath = strings.TrimSuffix(cleanPath, "/")
	cleanPath = strings.TrimSuffix(cleanPath, "\\")
	return &Assignment{
		Path:    cleanPath,
		Subject: subject,
		Course:  course,
		Weight:  weight,
	}, nil
}

func (a *Assignment) IsEquivalent(otherAssignment *Assignment) bool {
	if a.Path != otherAssignment.Path {
		return false
	}

	if a.Subject != otherAssignment.Subject {
		return false
	}

	if a.Course != otherAssignment.Course {
		return false
	}

	if a.Weight != otherAssignment.Weight {
		return false
	}

	if a.TotalTestCase != "" && a.TotalTestCase != otherAssignment.TotalTestCase {
		return false
	}

	return true
}
