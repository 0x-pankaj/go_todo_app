package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Todo      string `json:"todo"`
	User      *User  `json:"user"`
	Completed bool   `json:"completed"`
}

type User struct {
	FullName string `json:"fullname"`
}

var Todos []Todo

//middleware

func (c *Todo) isEmpty() bool {
	return c.Todo == ""
}

func main() {
	fmt.Println("todo app initilization")

	r := mux.NewRouter()

	r.HandleFunc("/", homeServe).Methods("GET")
	r.HandleFunc("/todos", getAllTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", getOneTodo).Methods("GET")
	r.HandleFunc("/todo/{user}", getUserTodo).Methods("GET")
	r.HandleFunc("/todo", createTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func homeServe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Router init")
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("<h1> Welcome To Todo App </h1> "))

}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting todos")
}

func getOneTodo(w http.ResponseWriter, r *http.Request) {

}

func getUserTodo(w http.ResponseWriter, r *http.Request) {

}

func createTodo(w http.ResponseWriter, r *http.Request) {

}

func updateTodo(w http.ResponseWriter, r *http.Request) {

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

}
