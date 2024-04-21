// Description: This file contains the controller functions for the image model.
// The controller functions are responsible for handling incoming HTTP requests
// and returning appropriate responses. The controller functions call the service
// functions to interact with the database and perform the necessary operations.
// The controller functions are defined in the controller package and are imported
// into the main file to be used by the server.
package controller

import (
	"log"
	"net/http"

	entity "models/entity"
	"service"

	"github.com/gin-gonic/gin"
)

/*
* GET /all
* Get all images
* @return {JSON} - entity.ResponseModel
* GetAllImages function is responsible for handling the GET /all request.
* It calls the GetAllImages service function to get all images from the database.
* If the service function returns an error, the controller function returns an error response.
* If the service function returns the images successfully, the controller function returns a success response.
 */
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

/*
* GET /id/:id
* Get image by id
* @param {string} - id
* @return {JSON} - entity.ResponseModel
* GetImageById function is responsible for handling the GET /id/:id request.
* It calls the GetImageById service function to get an image by id from the database.
* If the service function returns an error, the controller function returns an error response.
* If the service function returns the image successfully, the controller function returns a success response.
 */
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

/*
* GET /category/:category_id
* Get images by category
* @param {string} - category_id
* @return {JSON} - entity.ResponseModel
* GetImagesByCategory function is responsible for handling the GET /category/:category_id request.
* It calls the GetImagesByCategory service function to get images by category from the database.
* If the service function returns an error, the controller function returns an error response.
* If the service function returns the images successfully, the controller function returns a success response.
 */
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

/*
* GET /status/:status
* Get images by status
* @param {string} - status
* @return {JSON} - entity.ResponseModel
* GetImagesByStatus function is responsible for handling the GET /status/:status request.
* It calls the GetImagesByStatus service function to get images by status from the database.
* If the service function returns an error, the controller function returns an error response.
* If the service function returns the images successfully, the controller function returns a success response.
 */
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

/*
* GET /description/:description
* Get images by description
* @param {string} - description
* @return {JSON} - entity.ResponseModel
* GetImagesByDescription function is responsible for handling the GET /description/:description request.
* It calls the GetImagesByDescription service function to get images by description from the database.
* If the service function returns an error, the controller function returns an error response.
* If the service function returns the images successfully, the controller function returns a success response.
 */
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

/*
* POST /add
* Create image
* @param {JSON} - entity.Image
* @return {JSON} - entity.ResponseModel
* CreateImage function is responsible for handling the POST /add request.
* It binds the request body to an entity.Image struct and calls the CreateImage service function
* to create an image in the database. If the service function returns an error, the controller
* function returns an error response. If the service function returns the image successfully, the
* controller function returns a success response.
 */
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
