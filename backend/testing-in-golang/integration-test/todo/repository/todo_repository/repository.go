package todo_repository

import (
	"database/sql"
	"fmt"
	"log"
	"todo/utils"

	_ "github.com/go-sql-driver/mysql"
)

var (
	TodoRepo todoRepoInterface = &todoRepo{}
)

const (
	queryGetTodo     = "SELECT id, title, body, created_at FROM todos WHERE id=?;"
	queryInsertTodo  = "INSERT INTO todos(title, body, created_at) VALUES(?, ?, ?);"
	queryUpdateTodo  = "UPDATE todos SET title=?, body=? WHERE id=?;"
	queryDeleteTodo  = "DELETE FROM todos WHERE id=?;"
	queryGetAllTodos = "SELECT id, title, body, created_at FROM todos;"
)

type todoRepoInterface interface {
	Get(int64) (*Todo, utils.TodoErr)
	Create(*Todo) (*Todo, utils.TodoErr)
	Delete(int64) utils.TodoErr
	GetAll() ([]Todo, utils.TodoErr)
	Initialize(string, string, string, string, string, string) *sql.DB
}
type todoRepo struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) todoRepoInterface {
	return &todoRepo{db: db}
}

func (mr *todoRepo) Get(todoId int64) (*Todo, utils.TodoErr) {
	stmt, err := mr.db.Prepare(queryGetTodo)
	if err != nil {
		return nil, utils.NewInternalServerError(fmt.Sprintf("Error when trying to prepare todo: %s", err.Error()))
	}
	defer stmt.Close()

	var msg Todo
	result := stmt.QueryRow(todoId)
	if getError := result.Scan(&msg.Id, &msg.Title, &msg.Body, &msg.CreatedAt); getError != nil {
		fmt.Println("this is the error man: ", getError)
		return nil, utils.ParseError(getError)
	}
	return &msg, nil
}

func (mr *todoRepo) GetAll() ([]Todo, utils.TodoErr) {
	stmt, err := mr.db.Prepare(queryGetAllTodos)
	if err != nil {
		return nil, utils.NewInternalServerError(fmt.Sprintf("Error when trying to prepare all todos: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, utils.ParseError(err)
	}
	defer rows.Close()

	results := make([]Todo, 0)

	for rows.Next() {
		var msg Todo
		if getError := rows.Scan(&msg.Id, &msg.Title, &msg.Body, &msg.CreatedAt); getError != nil {
			return nil, utils.NewInternalServerError(fmt.Sprintf("Error when trying to get todo: %s", getError.Error()))
		}
		results = append(results, msg)
	}
	if len(results) == 0 {
		return nil, utils.NewNotFoundError("no records found")
	}
	return results, nil
}

func (mr *todoRepo) Create(msg *Todo) (*Todo, utils.TodoErr) {
	fmt.Println("WE REACHED THE DOMAIN")
	stmt, err := mr.db.Prepare(queryInsertTodo)
	if err != nil {
		return nil, utils.NewInternalServerError(fmt.Sprintf("error when trying to prepare user to save: %s", err.Error()))
	}
	fmt.Println("WE DIDNT REACH HERE")

	defer stmt.Close()

	insertResult, createErr := stmt.Exec(msg.Title, msg.Body, msg.CreatedAt)
	if createErr != nil {
		return nil, utils.ParseError(createErr)
	}
	msgId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, utils.NewInternalServerError(fmt.Sprintf("error when trying to save todo: %s", err.Error()))
	}
	msg.Id = msgId

	return msg, nil
}

func (mr *todoRepo) Delete(msgId int64) utils.TodoErr {
	stmt, err := mr.db.Prepare(queryDeleteTodo)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("error when trying to delete todo: %s", err.Error()))
	}
	defer stmt.Close()

	if _, err := stmt.Exec(msgId); err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("error when trying to delete todo %s", err.Error()))
	}
	return nil
}

func (mr *todoRepo) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	mr.db, err = sql.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal("This is the error connecting to the database:", err)
	}
	fmt.Printf("We are connected to the %s database", Dbdriver)

	return mr.db
}
