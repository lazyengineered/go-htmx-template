package main

import (
    "log/slog"
    "net/http"
)

func main() {
    s := http.NewServeMux()
    s.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world"))
    })

    slog.Info("Listening on http://localhost:3000...")

    http.ListenAndServe(":3000", s)
}
