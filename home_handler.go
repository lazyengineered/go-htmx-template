package main

import "net/http"

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
    return &HomeHandler{}
}

func (*HomeHandler) Home() http.HandlerFunc {
    t := NewTemplate("layout", "home")

    return func(w http.ResponseWriter, r *http.Request) {
        t.Execute(w, nil)
    }
}

func (*HomeHandler) Clicked() http.HandlerFunc {
    t := NewTemplate("home")

    return func(w http.ResponseWriter, r *http.Request) {
        t.ExecuteTemplate(w, "clicked", nil)
    }
}
