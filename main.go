package main

import (
	"log"

	"github.com/adjsky/gitjika/http/server"
)

func main() {
	server := server.New()
	log.Fatal(server.Listen())
}
