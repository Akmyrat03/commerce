package handler

import (
	"e-commerce/internal/categories/model"
	"e-commerce/internal/categories/service"
	handler "e-commerce/pkg/response"
	"net/http"
	"os"
	"path/filepath"
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
	name := c.PostForm("name")

	image, err := c.FormFile("image")
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uploadDir := "./uploads/categories"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	filePath := filepath.Join(uploadDir, image.Filename)

	if err := c.SaveUploadedFile(image, filePath); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	input := model.Category{
		Name:  name,
		Image: filePath,
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
	// Parse and validate category ID from URL
	idInput := c.Param("id")
	id, err := strconv.Atoi(idInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Get and validate the new name
	name := c.PostForm("name")
	if name == "" {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Category name is required")
		return
	}

	// Fetch existing category to get the current image path
	existingCategory, err := h.service.GetCategoryByID(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Error fetching category")
		return
	}

	// Handle the new image upload
	image, err := c.FormFile("image")
	var filePath string
	if err == nil {
		// Define upload directory and new image path
		uploadDir := "./uploads/categories"
		filePath = filepath.Join(uploadDir, image.Filename)

		// Delete the old image file if it exists
		if existingCategory.Image != "" {
			if err := os.Remove(existingCategory.Image); err != nil && !os.IsNotExist(err) {
				handler.NewErrorResponse(c, http.StatusInternalServerError, "Error deleting old image")
				return
			}
		}

		// Save the new image file
		if err := c.SaveUploadedFile(image, filePath); err != nil {
			handler.NewErrorResponse(c, http.StatusInternalServerError, "Error saving new image")
			return
		}
	} else {
		// No new image uploaded, keep the existing one
		filePath = existingCategory.Image
	}

	// Update the category in the database
	err = h.service.Update(id, name, filePath)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Respond with success message
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
