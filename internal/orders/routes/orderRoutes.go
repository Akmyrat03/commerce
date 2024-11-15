package routes

import (
	"e-commerce/internal/orders/handler"
	"e-commerce/internal/orders/repository"
	"e-commerce/internal/orders/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitOrderRoutes(DB *sqlx.DB, router *gin.RouterGroup) {
	orderRepo := repository.NewOrderRepository(DB)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	orderRoutes := router.Group("/orders")

	orderRoutes.POST("/add-order", orderHandler.CreateOrder)
	orderRoutes.GET("/view-all", orderHandler.GetOrders)
}
