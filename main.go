package main

import (
	"context"
	"fmt"

	controllers "controller"
	"data"
	messaging "messaging"
	"service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var DB *gorm.DB
	db := data.Init(DB)
	service.SetDB(db)
	restServer := NewServer()

	ctx := context.Background()
	go restServer.Run(":8080")
	go messaging.Connect(ctx)

	select {}
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
	s.router.GET("/api/v1/category/status/:status", controllers.GetCategoriesByStatus)
	s.router.POST("/api/v1/category/add", controllers.CreateCategory)
	s.router.PUT("/api/v1/category/update", controllers.UpdateCategory)
	s.router.GET("/api/v1/images/all", controllers.GetAllImages)
	s.router.GET("/api/v1/images/id/:id", controllers.GetImageById)
	s.router.GET("/api/v1/images/category/:category_id", controllers.GetImagesByCategory)
	s.router.GET("/api/v1/images/status/:status", controllers.GetImagesByStatus)
	s.router.GET("/api/v1/images/description/:description", controllers.GetImagesByDescription)
	s.router.POST("/api/v1/images/add", controllers.CreateImage)
	// s.router.DELETE("/api/v1/category/delete", controllers.DeleteCategory)
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
