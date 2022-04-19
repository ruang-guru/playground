package todo_repository

import (
	"strings"
	"time"
	"todo/utils"
)

type Todo struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *Todo) Validate() utils.TodoErr {
	m.Title = strings.TrimSpace(m.Title)
	m.Body = strings.TrimSpace(m.Body)
	if m.Title == "" {
		return utils.NewUnprocessibleEntityError("Please enter a valid title")
	}
	if m.Body == "" {
		return utils.NewUnprocessibleEntityError("Please enter a valid body")
	}
	return nil
}
