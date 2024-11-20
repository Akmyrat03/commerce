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

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product model.Product

	product.Name = c.PostForm("name")
	product.Description = c.PostForm("description")
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CategoryID, _ = strconv.Atoi(c.PostForm("category_id"))
	product.Status = c.PostForm("status")

	product.Images = []string{}
	form, err := c.MultipartForm()
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid form data")
		return
	}

	files := form.File["images"]

	uploadDir := "./uploads/products"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}

	for _, file := range files {

		filePath := filepath.Join(uploadDir, file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to upload image")
			return
		}

		product.Images = append(product.Images, filePath)
	}

	productID, err := h.service.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Product created successfully",
		"productID": productID,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idInput := c.Param("id")
	id, err := strconv.Atoi(idInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// existingProduct, err := h.service.GetProductByID(id)
	// if err != nil {
	// 	handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// fmt.Println("Image path before deletion:", existingProduct.Image)

	if err := h.service.DeleteProduct(id); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	// if existingProduct.Image != "" {
	// 	if err := os.Remove(existingProduct.Image); err != nil && !os.IsNotExist(err) {
	// 		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to delete image file")
	// 		return
	// 	}
	// }

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
