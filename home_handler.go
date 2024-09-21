package main

import (
    "html/template"
    "net/http"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
    return &HomeHandler{}
}

func (*HomeHandler) Home() http.HandlerFunc {
    t := template.Must(template.ParseFS(templatesFS, "templates/layout.html", "templates/home.html"))

    return func(w http.ResponseWriter, r *http.Request) {
        t.Execute(w, nil)
    }
}

func (*HomeHandler) Clicked() http.HandlerFunc {
    t := template.Must(template.ParseFS(templatesFS, "templates/home.html"))

    return func(w http.ResponseWriter, r *http.Request) {
        t.ExecuteTemplate(w, "clicked", nil)
    }
}
