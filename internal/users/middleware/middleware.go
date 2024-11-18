package middleware

import (
	"context"
	"e-commerce/internal/users/model"
	"e-commerce/internal/users/repository"
	"e-commerce/internal/users/service"
	"net/http"
	"strings"
	"time"

	handler "e-commerce/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type UserMiddleware struct {
	repo    *repository.UserRepository
	service *service.UserService
	redis   *redis.Client
}

func NewUserMiddleware(repo *repository.UserRepository, service *service.UserService, redis *redis.Client) *UserMiddleware {
	return &UserMiddleware{
		repo:    repo,
		service: service,
		redis:   redis,
	}
}

func (m *UserMiddleware) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if input.Username == "" || input.PhoneNumber == "" || input.Password == "" {
			handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
			c.Abort()
			return
		}

		validationMessage := service.ValidatePhoneNumber(input.PhoneNumber)
		if validationMessage != "Valid phone number." {
			handler.NewErrorResponse(c, http.StatusBadRequest, validationMessage)
			c.Abort()
			return
		}

		if len(input.Password) < 4 {
			handler.NewErrorResponse(c, http.StatusBadRequest, "Password must be at least 4 characters")
			c.Abort()
			return
		}

		// Check for existing user by username
		existingUser, err := m.repo.GetUserByField("username", input.Username)
		if err == nil && existingUser.Username != "" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Username already exists",
			})
			c.Abort()
			return
		}

		// Check for existing user by phone number
		existingNumberUser, err := m.repo.GetUserByField("phone_number", input.PhoneNumber)
		if err == nil && existingNumberUser.PhoneNumber != "" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Phone number already exists",
			})
			c.Abort()
			return
		}

		// Create the user since it doesn't exist
		if _, err := m.service.CreateUser(&input); err != nil {
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

func (m *UserMiddleware) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if input.Username == "" || input.Password == "" {
			handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
			return
		}

		token, err := m.service.GenerateToken(input.Username, input.Password)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		user, err := m.service.GetUser(input.Username, service.GeneratePasswordHash(input.Password))
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, "username or password is incorrect")
			return
		}

		if user.Role == "admin" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"token":    token,
				"redirect": "/admin/dashboard",
			})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"token":    token,
				"redirect": "/user/profile",
			})
		}
	}
}

func (m *UserMiddleware) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input format",
			})
			c.Abort()
			return
		}

		if input.Username == "" || input.Password == "" {
			handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
			return
		}

		err := m.service.SignOut(input.Username, input.Password)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, "username or password is incorrect")
			return
		}

		token := c.Request.Header.Get("Authorization")
		if token != "" {
			ctx := context.Background()
			m.redis.Set(ctx, "blacklist:"+token, "true", time.Hour*24)
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Accounted deleted successfully",
		})

	}
}

func (m *UserMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			handler.NewErrorResponse(c, http.StatusUnauthorized, "Token gereklidir")
			c.Abort()
			return
		}

		ctx := context.Background()

		blacklisted, err := m.redis.Get(ctx, "blacklist:"+token).Result()
		if err == nil && blacklisted == "true" {
			handler.NewErrorResponse(c, http.StatusUnauthorized, "Token gecersiz")
			c.Abort()
			return
		}

		c.Next()
	}
}

func (m *UserMiddleware) Profile() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			handler.NewErrorResponse(c, http.StatusUnauthorized, "Token gereklidir")
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		username, err := m.service.ValidateToken(token)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, "Geçersiz token")
			c.Abort()
			return
		}

		user, err := m.repo.GetUserByField("username", username)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusNotFound, "Kullanıcı bulunamadı")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username":     user.Username,
			"phone_number": user.PhoneNumber,
		})
	}
}

func (m *UserMiddleware) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := m.service.GetAll()
		if err != nil {
			handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
