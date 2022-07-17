package integration__tests

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"
	domain "todo/repository/todo_repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	queryTruncateTodo = "TRUNCATE TABLE Todos;"
	queryInsertTodo   = "INSERT INTO Todos(title, body, created_at) VALUES(?, ?, ?);"
	queryGetAllTodos  = "SELECT id, title, body, created_at FROM Todos;"
)

var (
	dbConn *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func database() {
	dbDriver := os.Getenv("DBDRIVER_TEST")
	username := os.Getenv("USERNAME_TEST")
	password := os.Getenv("PASSWORD_TEST")
	host := os.Getenv("HOST_TEST")
	database := os.Getenv("DATABASE_TEST")
	port := os.Getenv("PORT_TEST")

	dbConn = domain.TodoRepo.Initialize(dbDriver, username, password, port, host, database)
}

func refreshTodosTable() error {

	stmt, err := dbConn.Prepare(queryTruncateTodo)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error truncating Todos table: %s", err)
	}
	return nil
}

func seedOneTodo() (domain.Todo, error) {
	msg := domain.Todo{
		Title:     "the title",
		Body:      "the body",
		CreatedAt: time.Now(),
	}
	stmt, err := dbConn.Prepare(queryInsertTodo)
	if err != nil {
		panic(err.Error())
	}
	insertResult, createErr := stmt.Exec(msg.Title, msg.Body, msg.CreatedAt)
	if createErr != nil {
		log.Fatalf("Error creating Todo: %s", createErr)
	}
	msgId, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("Error creating Todo: %s", createErr)
	}
	msg.Id = msgId
	return msg, nil
}

func seedTodos() ([]domain.Todo, error) {
	todos := []domain.Todo{
		{
			Title:     "first title",
			Body:      "first body",
			CreatedAt: time.Now(),
		},
		{
			Title:     "second title",
			Body:      "second body",
			CreatedAt: time.Now(),
		},
	}
	stmt, err := dbConn.Prepare(queryInsertTodo)
	if err != nil {
		panic(err.Error())
	}
	for i, _ := range todos {
		_, createErr := stmt.Exec(todos[i].Title, todos[i].Body, todos[i].CreatedAt)
		if createErr != nil {
			return nil, createErr
		}
	}
	get_stmt, err := dbConn.Prepare(queryGetAllTodos)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := get_stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]domain.Todo, 0)

	for rows.Next() {
		var msg domain.Todo
		if getError := rows.Scan(&msg.Id, &msg.Title, &msg.Body, &msg.CreatedAt); getError != nil {
			return nil, err
		}
		results = append(results, msg)
	}
	return results, nil
}
