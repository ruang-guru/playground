package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todo/repository/todo_repository"
)

var (
	router = gin.Default()
)

func init() {
	//loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func StartApp() {

	dbdriver := os.Getenv("DBDRIVER")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	port := os.Getenv("PORT")

	// Init db
	todo_repository.TodoRepo.Initialize(dbdriver, username, password, port, host, database)
	//todo_repository.NewTodoRepository(db)

	fmt.Println("\nDATABASE STARTED")

	routes()

	router.Run(":8080")
}
