package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"service"
)

// GET /all
// Get all categories
func GetAllCategories(c *gin.Context) {
	res := service.GetAllImages()
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
