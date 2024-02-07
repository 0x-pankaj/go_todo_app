package main

import (
	"fmt"
	"net/http"
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

func main() {
	fmt.Println("todo app initilization")
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting todos")
}

func getOneTodo(w http.ResponseWriter, r *http.Request) {

}

func createTodo(w http.ResponseWriter, r *http.Request) {

}

func updateTodo(w http.ResponseWriter, r *http.Request) {

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

}
