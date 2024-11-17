package handler

import (
	"e-commerce/internal/products/model"
	"e-commerce/internal/products/service"
	handler "e-commerce/pkg/response"
	"fmt"
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

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")

	priceInput := c.PostForm("price")
	price, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid price format")
		return
	}

	catIdInput := c.PostForm("category_id")
	catID, err := strconv.Atoi(catIdInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	status := c.PostForm("status")

	file, err := c.FormFile("image")
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uploadDir := "./uploads/products"

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
		Image:       filePath,
		Status:      status,
		CategoryID:  catID,
	}

	if err := h.service.AddProduct(product); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idInput := c.Param("id")
	id, err := strconv.Atoi(idInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	existingProduct, err := h.service.GetProductByID(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Image path before deletion:", existingProduct.Image)

	if err := h.service.DeleteProduct(id); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	if existingProduct.Image != "" {
		if err := os.Remove(existingProduct.Image); err != nil && !os.IsNotExist(err) {
			handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to delete image file")
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *ProductHandler) GetAllPublishedProducts(c *gin.Context) {
	products, err := h.service.GetAllPublishedProducts()
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *ProductHandler) GetProductByCategoryName(c *gin.Context) {
	categoryName := c.Param("name")

	products, err := h.service.GetProductByCatName(categoryName)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *ProductHandler) LikeProduct(c *gin.Context) {
	var req model.LikedProduct
	if err := c.BindJSON(&req); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.LikeProduct(req.UserID, req.ProductID); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to like product")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product liked successfully",
	})
}

func (h *ProductHandler) UnlikeProduct(c *gin.Context) {
	var req model.LikedProduct
	if err := c.BindJSON(&req); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UnlikeProduct(req.UserID, req.ProductID); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product unliked successfully",
	})
}

func (h *ProductHandler) GetLikedProducts(c *gin.Context) {
	input := c.Query("user_id")
	userID, err := strconv.Atoi(input)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	products, err := h.service.GetLikedProducts(userID)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to fetch liked products")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"like_products": products,
	})
}
