package main

import (
	"html/template"
	"net/http"
)

// Todo -
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData -
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// HomeHandler -
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	var tmpl = template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8181", nil)
}
