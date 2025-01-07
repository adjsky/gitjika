package main

import (
	"github.com/adjsky/gitjika/internal/config"
	"github.com/adjsky/gitjika/server"
)

func main() {
	cfg := config.New()
	server := server.New(cfg)
	panic(server.Listen())
}
