package handler

import (
	"e-commerce/internal/orders/model"
	"e-commerce/internal/orders/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		UserID int               `json:"user_id"`
		Items  []model.OrderItem `json:"items"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	order := model.Order{
		UserID:     req.UserID,
		TotalPrice: calculateTotalPrice(req.Items),
		Status:     "pending",
	}

	if err := h.service.CreateOrder(&order, req.Items); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	userIDInput := c.Query("user_id")
	userID, err := strconv.Atoi(userIDInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
	}
	orders, err := h.service.GetOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func calculateTotalPrice(items []model.OrderItem) float64 {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}
