package controller

import (
	"log"
	"net/http"

	entity "models/entity"
	"service"

	"github.com/gin-gonic/gin"
)

// GET /all
// Get all images
func GetAllImages(c *gin.Context) {
	res := service.GetAllImages()
	if res.Error != nil {
		if res.Error == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"data": res})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetImageById(c *gin.Context) {
	id := c.Param("id")

	res := service.GetImageById(id)
	log.Printf("error: %v", res.Error)
	if res.Error != nil {
		if res.Error == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"data": res})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetImagesByCategory(c *gin.Context) {
	id := c.Param("category_id")

	res := service.GetImagesByCategory(id)
	if res.Error != nil {
		if res.Error == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"data": res})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetImagesByStatus(c *gin.Context) {
	status := c.Param("status")

	res := service.GetImagesByStatus(status)
	if res.Error != nil {
		if res.Error == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"data": res})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetImagesByDescription(c *gin.Context) {
	description := c.Param("description")

	res := service.GetImagesByDescription(description)
	if res.Error != nil {
		if res.Error == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"data": res})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func CreateImage(c *gin.Context) {
	var image entity.Image
	if c.Request.ContentLength == 0 {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"data": entity.SetResponse(nil, "Empty request body", "error")},
		)
		return
	}
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": entity.SetResponse(nil, err.Error(), "error")})
		return
	}

	res := service.CreateImage(image)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
