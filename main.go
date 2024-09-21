package main

import (
    "embed"
    "html/template"
    "log/slog"
    "net/http"
)

var (
    //go:embed templates/*
    templatesFS embed.FS
)

func main() {
    s := http.NewServeMux()

    t := template.Must(template.ParseFS(templatesFS, "templates/layout.html", "templates/home.html"))
    s.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        t.Execute(w, nil)
    })

    s.HandleFunc("POST /clicked", func(w http.ResponseWriter, r *http.Request) {
        t.ExecuteTemplate(w, "clicked", nil)
    })

    slog.Info("Listening on http://localhost:3000...")

    http.ListenAndServe(":3000", s)
}
