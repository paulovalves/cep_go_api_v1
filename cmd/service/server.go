package service

import (
	"data"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func newServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Run(addr string) {
	r := s.router.Run(addr)
	if r != nil {
		log.Fatalf("Error while running server: %v", r)
	}
}

func Start(addr string) {
	data.Init()
	server := newServer()
	server.Run(addr)
}
