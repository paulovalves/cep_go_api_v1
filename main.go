// Description: This file is the entry point of the application. It initializes the database connection, starts the REST server and the messaging service.
// The main function initializes the database connection, sets the database instance in the service package, creates a new REST server, and starts the server.
// The main function also starts the messaging service in a separate goroutine.
package main

import (
	"context"
	"data"
	"fmt"
	"service"

	controllers "controller"

	messaging "messaging"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
* The main function initializes the database connection, sets the database instance in the service package,
* creates a new REST server, and starts the server. The main function also starts the messaging service in a separate goroutine.
 */
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

/*
* NewServer function creates a new REST server instance.
* Return: [Server] - REST server instance
 */
func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

/*
* Server struct represents a REST server.
* It has a router field that is an instance of the gin.Engine.
 */
type Server struct {
	router *gin.Engine
}

/*
* Run function starts the REST server on the specified address.
* Param: addr - server address
 */
func (s *Server) Run(addr string) {
	// Category routes
	s.router.GET("/api/v1/category/id/:id", controllers.GetCategoryById)
	s.router.GET("/api/v1/category/all", controllers.GetCategories)
	s.router.GET("/api/v1/category/status/:status", controllers.GetCategoriesByStatus)
	s.router.POST("/api/v1/category/add", controllers.CreateCategory)
	s.router.PUT("/api/v1/category/update", controllers.UpdateCategory)
	// s.router.DELETE("/api/v1/category/delete", controllers.DeleteCategory)

	// Image routes
	s.router.GET("/api/v1/images/all", controllers.GetAllImages)
	s.router.GET("/api/v1/images/id/:id", controllers.GetImageById)
	s.router.GET("/api/v1/images/category/:category_id", controllers.GetImagesByCategory)
	s.router.GET("/api/v1/images/status/:status", controllers.GetImagesByStatus)
	s.router.GET("/api/v1/images/description/:description", controllers.GetImagesByDescription)
	s.router.POST("/api/v1/images/add", controllers.CreateImage)
	s.router.PUT("/api/v1/images/update", controllers.UpdateImage)
	s.router.DELETE("/api/v1/images/id/:id", controllers.DeleteImage)
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
