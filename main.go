package main

import (
	"log"

	"github.com/adjsky/gitjika/internal/http/server"
)

func main() {
	server := server.New()
	log.Fatal(server.Listen())
}
