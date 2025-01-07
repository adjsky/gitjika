package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adjsky/gitjika/internal/config"
	"github.com/adjsky/gitjika/server/handlers"
)

type Server struct {
	mux *http.ServeMux
	cfg config.Config
}

func New(cfg config.Config) Server {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handlers.Index))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return Server{mux, cfg}
}

func (s Server) Listen() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port)

	slog.Info("Start listening", slog.String("addr", addr))

	return http.ListenAndServe(addr, s.mux)
}
