package routes

import (
	"e-commerce/internal/cart_items/handler"
	"e-commerce/internal/cart_items/repository"
	"e-commerce/internal/cart_items/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitCartRoutes(db *sqlx.DB, router *gin.RouterGroup) {
	cartItemRepo := repository.NewCartItemRepository(db)
	cartItemServ := service.NewCartItemService(cartItemRepo)
	cartItemHand := handler.NewCartItemHandler(cartItemServ)

	cartItemRoutes := router.Group("item")

	cartItemRoutes.POST("/add-cart-item", cartItemHand.CreateCartItem)
}
