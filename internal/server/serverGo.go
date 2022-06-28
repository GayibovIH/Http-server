package server

import (
	"github.com/gayibovih/http-server/internal/handler/httpHandler"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
	Http   *httpHandler.HttpHandler
}

func New() *Server {
	r := chi.NewRouter()
	handle := httpHandler.New()
	r.Get("/fib/{num}", handle.GetHandle)
	r.Post("/fib/{num}", handle.CalculateHandle)
	return &Server{Router: r, Http: handle}

}
