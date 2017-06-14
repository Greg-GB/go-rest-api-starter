package main

import (
	"net/http"

	h "github.com/greg-gb/go-rest-api-starter/handlers"
	m "github.com/greg-gb/go-rest-api-starter/middleware"
)

func main() {
	r := http.NewServeMux()
	r.Handle("/", m.Wrapper(h.Root(), m.AddRequestID))
	r.Handle("/auth", m.Wrapper(h.Auth(), m.CheckAuth, m.AddRequestID))

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	s.ListenAndServe()
}
