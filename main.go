// Description: This file is the entry point of the application. It initializes the database connection, starts the REST server and the messaging service.
// The main function initializes the database connection, sets the database instance in the service package, creates a new REST server, and starts the server.
// The main function also starts the messaging service in a separate goroutine.
package main

import (
	"data"
	"log"
	"service"

	controllers "controller"

	middleware "middleware"

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

	// ctx := context.Background()
	go restServer.Run(":8080")
	// go messaging.Connect(ctx)

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
	g := gin.Default()
	routes := g.Group("/api/v1")
	{
		publicRoutes := routes.Group("/public")
		{
			log.Printf("Public routes: %v", publicRoutes)
			// Category routes
			publicRoutes.GET("/category/id/:id", controllers.GetCategoryById)
			publicRoutes.GET("/category/all", controllers.GetCategories)
			publicRoutes.GET("/category/status/:status", controllers.GetCategoriesByStatus)
			publicRoutes.GET("/images/all", controllers.GetAllImages)
			publicRoutes.GET("/images/id/:id", controllers.GetImageById)
			publicRoutes.GET("/images/category/:category_id", controllers.GetImagesByCategory)
			publicRoutes.GET("/images/status/:status", controllers.GetImagesByStatus)
			publicRoutes.GET("/images/description/:description", controllers.GetImagesByDescription)
			publicRoutes.POST("/auth/register", controllers.Register)
			publicRoutes.POST("/auth/login", controllers.Login)
		}

		protectedRoutes := routes.Group("/protected")
		{
			protectedRoutes.POST("/category/add", controllers.CreateCategory)
			protectedRoutes.PUT("/category/update", controllers.UpdateCategory)
			// s.router.DELETE("/api/v1/category/delete", controllers.DeleteCategory)

			// Image routes
			protectedRoutes.POST("/images/add", controllers.CreateImage)
			protectedRoutes.PUT("/images/update", controllers.UpdateImage)
			protectedRoutes.DELETE("/images/id/:id", controllers.DeleteImage)
		}
		protectedRoutes.Use(middleware.AuthenticationMiddleware())
	}
	if err := g.Run(addr); err != nil {
		log.Println(err)
	}
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
