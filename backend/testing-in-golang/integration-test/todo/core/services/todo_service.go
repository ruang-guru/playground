package services

import (
	"time"
	todoRepo "todo/repository/todo_repository"
	"todo/utils"
)

var (
	TodosService todoServiceInterface = &todosService{}
)

type todosService struct{}

type todoServiceInterface interface {
	GetTodo(int64) (*todoRepo.Todo, utils.TodoErr)
	CreateTodo(*todoRepo.Todo) (*todoRepo.Todo, utils.TodoErr)
	DeleteTodo(int64) utils.TodoErr
	GetAllTodos() ([]todoRepo.Todo, utils.TodoErr)
}

func (m *todosService) GetTodo(msgId int64) (*todoRepo.Todo, utils.TodoErr) {
	todo, err := todoRepo.TodoRepo.Get(msgId)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (m *todosService) GetAllTodos() ([]todoRepo.Todo, utils.TodoErr) {
	todos, err := todoRepo.TodoRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (m *todosService) CreateTodo(todo *todoRepo.Todo) (*todoRepo.Todo, utils.TodoErr) {
	if err := todo.Validate(); err != nil {
		return nil, err
	}
	todo.CreatedAt = time.Now()
	todo, err := todoRepo.TodoRepo.Create(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (m *todosService) DeleteTodo(msgId int64) utils.TodoErr {
	msg, err := todoRepo.TodoRepo.Get(msgId)
	if err != nil {
		return err
	}
	deleteErr := todoRepo.TodoRepo.Delete(msg.Id)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
