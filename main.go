package main

import (
	"log"

	"github.com/adjsky/gitjika/server"
)

func main() {
	server := server.New()
	log.Fatal(server.Listen())
}
