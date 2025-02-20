package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hokamsingh/go-backend-template/internal/service"
)

// CreateUser handles user creation.
func CreateUser(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service.CreateUserInput

		// Bind JSON request body to input struct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the service to create user
		createdUser, err := userService.Create(c.Request.Context(), input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, createdUser)
	}
}

// GetUser handles fetching a user by ID.
func GetUser(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")

		// Convert id string to uint
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		// Call the service to get the user
		user, err := userService.GetByID(c.Request.Context(), uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
