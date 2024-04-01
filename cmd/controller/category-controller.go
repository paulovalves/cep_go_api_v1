package controller

import (
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

// GET /all
// Get all categories
func GetCategories(c *gin.Context) {
	data, err := service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /:id
// Get one category by id
func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	log.Println("c ID: ", id)
	data, err := service.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}