package main

import (
    "log/slog"
    "net/http"
)

func main() {
    s := http.NewServeMux()

    homeHandler := NewHomeHandler()

    s.HandleFunc("GET /", homeHandler.Home())
    s.HandleFunc("POST /clicked", homeHandler.Clicked())

    slog.Info("Listening on http://localhost:3000...")
    http.ListenAndServe(":3000", s)
}
