package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Todo struct {
	Todoid    string `json:"id"`
	Todo      string `json:"todo"`
	User      *User  `json:"user"`
	Completed bool   `json:"completed"`
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var Todos []Todo

//middleware

func (c *Todo) isEmpty() bool {
	return c.Todo == ""
}

func main() {
	fmt.Println("todo app initilization")

	r := mux.NewRouter()

	//seeding
	Todos = append(Todos, Todo{Todoid: "1", Todo: "have to complete golang", User: &User{FirstName: "pankaj", LastName: "singh"}, Completed: false})
	Todos = append(Todos, Todo{Todoid: "3", Todo: "Database", User: &User{"Rohan", "singh"}, Completed: true})
	Todos = append(Todos, Todo{"5", "Devops", &User{"Rohan", "singh"}, false})

	r.HandleFunc("/", homeServe).Methods("GET")
	r.HandleFunc("/todos", getAllTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", getOneTodo).Methods("GET")
	r.HandleFunc("/user/{user}", getUserTodo).Methods("GET")
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
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&Todos)
	if err != nil {
		http.Error(w, "Failed to decode Request Body", http.StatusBadRequest)
	}

}

func getOneTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "Id must be required", http.StatusBadRequest)
	}

	for _, t := range Todos {
		if t.Todoid == params["id"] {
			err := json.NewEncoder(w).Encode(t)
			if err != nil {
				http.Error(w, "Failed to decode todo", http.StatusBadRequest)
			}
			return
		}
	}
	http.Error(w, "Not found with given id ", http.StatusNotFound)

}

func getUserTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting user todo")
	params := mux.Vars(r)

	fmt.Println(params["user"])

	if params["user"] == "" {
		http.Error(w, "Must have user to search on user", http.StatusBadRequest)
	}

	var userTodo []Todo

	for _, u := range Todos {
		if strings.ToLower(u.User.FirstName) == strings.ToLower(params["user"]) {
			userTodo = append(userTodo, u)
		}
	}

	if len(userTodo) == 0 {
		http.Error(w, "Not found with given user", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(&userTodo)
	if err != nil {
		http.Error(w, "Failed to decode data into json", http.StatusBadRequest)
		return
	}

}

func createTodo(w http.ResponseWriter, r *http.Request) {

}

func updateTodo(w http.ResponseWriter, r *http.Request) {

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

}
