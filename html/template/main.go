package main

import (
	"net/http"

	"github.com/tayron/golang-estudos/html/template/bootstrap"
)

// Todo -
type Todo struct {
	Title string
	Done  bool
}

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	parameters := struct {
		PageTitle string
		Todos     []Todo
	}{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	bootstrap.LoadView(w, "homePage", parameters)
}

func main() {

	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8181", nil)
}
