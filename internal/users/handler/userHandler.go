package handler

import (
	"e-commerce/internal/users/model"
	"e-commerce/internal/users/service"
	handler "e-commerce/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserService(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// func (h *UserHandler) SignUp(c *gin.Context) {
// 	var input model.User
// 	if err := c.BindJSON(&input); err != nil {
// 		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	if input.Username == "" || input.Email == "" || input.Password == "" {
// 		handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
// 		return
// 	}

// 	if len(input.Password) < 4 {
// 		handler.NewErrorResponse(c, http.StatusBadRequest, "Password must be at least 4 characters")
// 		return
// 	}

// 	_, err := h.service.CreateUser(&input)
// 	if err != nil {
// 		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, map[string]interface{}{
// 		"username": input.Username,
// 		"password": input.Password,
// 	})
// }

func (h *UserHandler) Login(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "" || input.Password == "" {
		handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
		return
	}

	if len(input.Password) < 4 {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Password must be at least 4 characters")
		return
	}

	token, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.GetUser(input.Username, service.GeneratePasswordHash(input.Password))
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

func (h *UserHandler) SignOut(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "" || input.Password == "" {
		handler.NewErrorResponse(c, http.StatusBadRequest, "All fields are required")
		return
	}

	err := h.service.SignOut(input.Username, input.Password)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusUnauthorized, "username or password is incorrect")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Accounted deleted successfully",
	})
}
