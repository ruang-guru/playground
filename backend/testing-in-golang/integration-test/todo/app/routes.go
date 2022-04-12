package app

import "todo/controllers"

func routes() {
	router.GET("/todos/:todo_id", controllers.GetTodo)
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.CreateTodo)
	//router.PUT("/todos/:todo_id", controllers.UpdateTodo)
	router.DELETE("/todos/:todo_id", controllers.DeleteTodo)
}
