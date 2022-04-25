package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8081"

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !AllowOnlyGET(w, r) {
		return
	}

	if !Auth(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{Id: "001", Name: "Gina", Grade: 2})
	students = append(students, &Student{Id: "002", Name: "Indri", Grade: 2})
	students = append(students, &Student{Id: "003", Name: "Dwi", Grade: 3})
}

const USERNAME = "aditira"
const PASSWORD = "aditira123"

// Middleware: Auth
func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}

// Test with username: batman, password: secret
//$ curl -X GET --user aditira:aditira123 http://localhost:8081/student
//$ curl -X GET --user aditira:aditira123 http://localhost:8081/student?id=002

// Generate base64 encoded string: batman:secret
// https://www.base64encode.org/

// Test with base64 encoded:
//$ curl -X GET -H "Authorization: Basic YWRpdGlyYTphZGl0aXJhMTIz" http://localhost:8081/student
//$ curl -X GET -H "Authorization: Basic YWRpdGlyYTphZGl0aXJhMTIz" http://localhost:8081/student?id=s002
