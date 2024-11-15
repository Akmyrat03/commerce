package handler

import (
	"e-commerce/internal/cart_items/model"
	"e-commerce/internal/cart_items/service"
	handler "e-commerce/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartItemHandler struct {
	service *service.CartItemService
}

func NewCartItemHandler(service *service.CartItemService) *CartItemHandler {
	return &CartItemHandler{service: service}
}

func (h *CartItemHandler) CreateCartItem(c *gin.Context) {
	var input model.CartItem
	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cartItem, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create cart item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cartItem": cartItem})
}

func (h *CartItemHandler) GetAllCartItems(c *gin.Context) {
	idInput := c.Param("id")
	id, err := strconv.Atoi(idInput)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.service.GetAll(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"cartItems": items})
}
