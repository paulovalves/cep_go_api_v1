package main

import (
	"log"
	"net/http"
	"time"

	entity "models/entity"
)

func main() {
	server := newServer()
	server.run(":8080")
}

func newServer() *entity.Server {
	return &entity.Server{
		router: gin.Default(),
	}
}

func (s *entity.Server) run(addr) {
	r := s.router.Run(addr)
	if r != nil {
		log.Fatalf("Error while running server: %v", r)
	}
}
