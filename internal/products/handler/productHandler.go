package handler

import (
	"e-commerce/internal/products/model"
	"e-commerce/internal/products/service"
	handler "e-commerce/pkg/response"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// func (h *ProductHandler) CreateProduct(c *gin.Context) {
// 	var product model.Product
// 	if err := c.BindJSON(&product); err != nil {
// 		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		handler.NewErrorResponse(c, http.StatusBadGateway, err.Error())
// 		return
// 	}

// 	uploadDir := "/uploads"

// 	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
// 		os.Mkdir(uploadDir, 0755)
// 	}

// 	filePath := filepath.Join(uploadDir, file.Filename)

// 	if err := c.SaveUploadedFile(file, filePath); err != nil {
// 		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed too upload file")
// 		return
// 	}

// 	product.Image = filePath

// 	if err := h.service.AddProduct(product); err != nil {
// 		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Product created successfully",
// 	})
// }

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")

	priceInput := c.PostForm("price")
	price, err := strconv.Atoi(priceInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid price format")
		return
	}

	catIDInput := c.PostForm("category_id")
	category_id, err := strconv.Atoi(catIDInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid category_id format")
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uploadDir := "./uploads"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	filePath := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	product := model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  category_id,
		Image:       filePath,
	}

	if err := h.service.AddProduct(product); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}
