package middleware

import (
	"e-commerce/internal/users/model"
	"e-commerce/internal/users/repository"
	"net/http"

	handler "e-commerce/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserMiddleware struct {
	repo *repository.UserRepository
}

func NewUserMiddleware(repo *repository.UserRepository) *UserMiddleware {
	return &UserMiddleware{repo: repo}
}

func (m *UserMiddleware) CheckUserExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input format",
			})
			c.Abort()
			return
		}

		if input.Username == "" || input.Email == "" || input.Password == "" {
			handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
			c.Abort()
			return
		}

		if len(input.Password) < 4 {
			handler.NewErrorResponse(c, http.StatusBadRequest, "Password must be at least 4 characters")
			c.Abort()
			return
		}

		// Check for existing user by username
		existingUser, err := m.repo.GetUserByUsername(input.Username)
		if err == nil && existingUser.Username != "" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Username already exists",
			})
			c.Abort()
			return
		}

		// Check for existing user by email
		existingEmailUser, err := m.repo.GetUserByEmail(input.Email)
		if err == nil && existingEmailUser.Email != "" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email already exists",
			})
			c.Abort()
			return
		}

		// Create the user since it doesn't exist
		if _, err := m.repo.CreateUser(&input); err != nil {
			handler.NewErrorResponse(c, http.StatusInternalServerError, "Could not create user")
			return
		}

		// Respond with a success message
		c.JSON(http.StatusCreated, gin.H{
			"message":  "User created successfully",
			"username": input.Username,
		})
		c.Next()
	}
}
