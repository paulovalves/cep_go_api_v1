package main

import (
	"data"
	"fmt"
	"service"

	controllers "controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var DB *gorm.DB
	db := data.Init(DB)
	service.SetDB(db)
	server := NewServer()
	server.Run(":8080")
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

type Server struct {
	router *gin.Engine
}

func (s *Server) Run(addr string) {
	s.router.GET("/api/v1/category/id/:id", controllers.GetCategoryById)
	s.router.GET("/api/v1/category/all", controllers.GetCategories)
	r := s.router.Run(addr)
	if r != nil {
		fmt.Println(r)
	}
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
