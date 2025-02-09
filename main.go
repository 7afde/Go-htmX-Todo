package main

import (
	"log"
	"net/http"
	"text/template"
)
type Todo struct {
	Id      int
	Message string
}

var data = map[string][]Todo{
	"Todos": {
		Todo{Id: 1, Message: "Buy Milk"},
	},
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))

	templ.Execute(w, data)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	templ := template.Must(template.ParseFiles("index.html"))
	todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
	data["Todos"] = append(data["Todos"], todo)

	templ.ExecuteTemplate(w, "todo-list-element", todo)
}

func main() {
	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
