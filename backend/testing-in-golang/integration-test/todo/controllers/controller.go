package controllers

import (
	"net/http"
	"strconv"
	"todo/core/services"
	todoRepo "todo/repository/todo_repository"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

//Since we are going for the todo id more than we, we extracted this functionality to a function so we can have a DRY code.
func getTodoId(todoIdParam string) (int64, utils.TodoErr) {
	todoId, todoErr := strconv.ParseInt(todoIdParam, 10, 64)
	if todoErr != nil {
		return 0, utils.NewBadRequestError("todo id should be a number")
	}
	return todoId, nil
}

func GetTodo(c *gin.Context) {
	todoId, err := getTodoId(c.Param("todo_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	todo, getErr := services.TodosService.GetTodo(todoId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func GetAllTodos(c *gin.Context) {
	todos, getErr := services.TodosService.GetAllTodos()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo *todoRepo.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		theErr := utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	todo, err := services.TodosService.CreateTodo(todo)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func DeleteTodo(c *gin.Context) {
	todoId, err := getTodoId(c.Param("todo_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	if err := services.TodosService.DeleteTodo(todoId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
