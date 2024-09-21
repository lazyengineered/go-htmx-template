package main

import (
    "embed"
    "log/slog"
    "net/http"
)

var (
    //go:embed templates/*.html templates/*/*.html
    templatesFS embed.FS
)

func main() {
    s := http.NewServeMux()

    homeHandler := NewHomeHandler()
    todoHandler := NewTodoHandler()

    s.HandleFunc("GET /", homeHandler.Home())
    s.HandleFunc("POST /clicked", homeHandler.Clicked())

    s.HandleFunc("GET /todos", todoHandler.Index())

    slog.Info("Listening on http://localhost:3000...")
    http.ListenAndServe(":3000", s)
}
