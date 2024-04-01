package server

import (
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

func Start() *Server {
	return newServer()
}
