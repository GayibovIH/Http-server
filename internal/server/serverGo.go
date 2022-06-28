package server

import "github.com/go-chi/chi/v5"

type Server struct {
	Router *chi.Mux
}

func New() *Server {
	r := chi.NewRouter()
	return &Server{Router: r}
}
