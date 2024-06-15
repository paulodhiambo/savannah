package main

import (
	"backend/internal/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
