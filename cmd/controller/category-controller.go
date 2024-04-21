package controller

import (
	"log"
	"net/http"

	entity "models/entity"
	"service"
	"utils"

	"github.com/gin-gonic/gin"
)

// GET /all
// Get all categories
func GetCategories(c *gin.Context) {
	res := service.GetCategories()
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// GET /:id
// Get one category by id
func GetCategoryById(c *gin.Context) {
	id := c.Param("id")

	if !utils.IsValidUUID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}
	res := service.GetCategoryById(id)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// GET /:status
// Get all categories by status
func GetCategoriesByStatus(c *gin.Context) {
	status := c.Param("status")

	res := service.GetCategoriesByStatus(status)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// POST
// Create new category
func CreateCategory(c *gin.Context) {
	var category entity.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": entity.SetResponse(nil, err.Error(), "error")})
	}

	res := service.CreateCategory(category)
	if res.Error != nil {
		log.Fatalf("Error while creating category: %v", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// UPDATE
// Update an existing category
func UpdateCategory(c *gin.Context) {
	var category entity.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res := service.UpdateCategory(category)
	if res.Error != nil {
		log.Fatalf("Error while updating category: %v", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
