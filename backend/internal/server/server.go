package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"gatcha/internal/config"
)

type Server struct {
	Config *config.Config
	Router *chi.Mux
}

func NewServer(cfg *config.Config, r *chi.Mux) *Server {
	return &Server{
		Config: cfg,
		Router: r,
	}
}

func (s *Server) Start() {
	log.Printf("Server started on port %s", s.Config.Server.Port)

	err := http.ListenAndServe(":"+s.Config.Server.Port, s.Router)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
