package handler

import (
	"e-commerce/internal/categories/model"
	"e-commerce/internal/categories/service"
	handler "e-commerce/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input model.Category
	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Category name is required")
		return
	}

	id, err := h.service.Create(&input)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *CategoryHandler) DeleteCategoryByID(c *gin.Context) {
	idInput := c.Param("id")
	id, err := strconv.Atoi(idInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}

func (h *CategoryHandler) UpdateCategoryByID(c *gin.Context) {
	idInput := c.Param("id")
	id, err := strconv.Atoi(idInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input model.Category
	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Category name is required")
		return
	}

	if input.Name == "" {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Category name is required")
		return
	}

	err = h.service.Update(id, input.Name)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category successfully updated",
	})
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.service.Get()
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}
