package controller

import (
	"auth"
	"net/http"

	entity "models/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"data": entity.SetResponse(nil, "Invalid request", "error")},
		)
		return
	}

	if user.Email == "user" && user.Password == "password" {
		token, err := auth.CreateToken(user.Email)
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"data": entity.SetResponse(nil, "Login failed", "error")})
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{"data": entity.SetResponse(token, nil, "success")},
		)
		return
	} else {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"data": entity.SetResponse(nil, "Invalid credentials", "error")},
		)
		return
	}
}

func Register(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"data": entity.SetResponse(nil, "Invalid request", "error")},
		)
		return
	}

	user.Id = uuid.New()
	user.Username = "user"
	user.Password = "password"
	user.Name = "User"
	user.Email = "email@email.com"
	user.Role = "user"

	c.JSON(
		http.StatusOK,
		gin.H{"data": entity.SetResponse(user, nil, "success")},
	)
}
