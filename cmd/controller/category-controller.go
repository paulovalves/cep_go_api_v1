package controller

import (
	"net/http"
	"service"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	data, err := service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
