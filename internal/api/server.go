package api

import (
	"log/slog"
	"net/http"

	"Regma/internal/config"
)

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
	mux    *http.ServeMux
}

func NewServer(cfg *config.Config, logger *slog.Logger) *Server {
	s := &Server{
		cfg:    cfg,
		logger: logger,
		mux:    http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /health", s.handleHealth)

	// TODO: Add your routes here
	// s.mux.HandleFunc("POST /api/v1/investors", s.handleCreateInvestor)
}

func (s *Server) Run() error {
	return http.ListenAndServe(":"+s.cfg.Port, s.mux)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
