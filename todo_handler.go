package main

import (
    "html/template"
    "net/http"
)

type TodoHandler struct {
}

func NewTodoHandler() *TodoHandler {
    return &TodoHandler{}
}

func (*TodoHandler) Index() http.HandlerFunc {
    t := template.Must(template.ParseFS(templatesFS, "templates/layout.html", "templates/todos/index.html"))

    return func(w http.ResponseWriter, r *http.Request) {
        t.Execute(w, nil)
    }
}
