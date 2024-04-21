// Description: This file includes all the functions that are used to handle the requests
// made to the category routes. The functions are responsible for calling the service
// functions and returning the response to the client. The functions include GetCategories,
// GetCategoryById, GetCategoriesByStatus, CreateCategory, and UpdateCategory. The GetCategories
// function returns all categories, GetCategoryById returns a category by its id, GetCategoriesByStatus
// returns all categories by status, CreateCategory creates a new category, and UpdateCategory updates
// an existing category. The functions use the service functions to interact with the database and
// perform the required operations.
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
/*
	GetCategories is a controller function that returns all categories
*/
func GetCategories(c *gin.Context) {
	res := service.GetCategories()
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

/*
* GET /:id
* Get one category by id
* GetCategoryById is a controller function that returns a category by its id
* @param c *gin.Context
* Api responses: 200 - OK, 400 - BAD REQUEST, 404 - NOT FOUND, 500 - INTERNAL SERVER ERROR
 */
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

/*
* GET /:status
* Get all categories by status
* GetCategoriesByStatus is a controller function that returns all categories by status (ativo, inativo ou removido)
* @param c *gin.Context
* Api responses: 200 - OK, 400 - BAD REQUEST, 404 - NOT FOUND, 500 - INTERNAL SERVER ERROR
 */
func GetCategoriesByStatus(c *gin.Context) {
	status := c.Param("status")

	res := service.GetCategoriesByStatus(status)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": res})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

/*
* POST
* Create new category
* CreateCategory is a controller function that creates a new category
* @param c *gin.Context
* Api responses: 200 - OK, 400 - BAD REQUEST, 500 - INTERNAL SERVER ERROR
 */
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

/*
* UPDATE
* Update an existing category
* UpdateCategory is a controller function that updates an existing category
* @param c *gin.Context
* Api responses: 200 - OK, 400 - BAD REQUEST, 500 - INTERNAL SERVER ERROR
 */
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
