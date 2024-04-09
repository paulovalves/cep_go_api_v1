package controller

import (
	"log"
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

// GET /all
// Get all images
func GetAllImages(c *gin.Context) {
	res := service.GetAllImages()
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetImageById(c *gin.Context) {
	id := c.Param("id")

	res := service.GetImageById(id)
	log.Printf("error: %v", res.Error)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
