package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func byebye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("byebye"))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello", welcome)
	r.Get("/bye", byebye)

	http.ListenAndServe(":3000", r)
}
