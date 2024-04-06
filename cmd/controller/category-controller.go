package controller

import (
	"log"
	"net/http"
	"service"
	"utils"

	"github.com/gin-gonic/gin"
	entity "models/entity"
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

	if !utils.IsValidUUID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	data, err := service.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /:status
// Get all categories by status
func GetCategoriesByStatus(c *gin.Context) {
	status := c.Param("status")

	data, err := service.GetCategoriesByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// POST
// Create new category
func CreateCategory(c *gin.Context) {
	var category entity.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := service.CreateCategory(category)
	if err != nil {
		log.Fatalf("Error while creating category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
