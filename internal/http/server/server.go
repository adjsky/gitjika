package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/adjsky/gitjika/internal/http/handlers"
)

type Server struct {
	mux *http.ServeMux
}

func New() Server {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handlers.Index))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return Server{mux}
}

func (s Server) Listen() error {
	addr := ":6969"

	slog.Info(fmt.Sprintf("Server listening on %v", addr))

	return http.ListenAndServe(addr, s.mux)
}
