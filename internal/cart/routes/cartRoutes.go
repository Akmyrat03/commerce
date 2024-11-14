package routes

import (
	"e-commerce/internal/cart/handler"
	"e-commerce/internal/cart/repository"
	"e-commerce/internal/cart/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitCartRoutes(db *sqlx.DB, router *gin.RouterGroup) {
	cartRepo := repository.NewCartRepository(db)
	cartServ := service.NewCartService(cartRepo)
	cartHand := handler.NewCartHandler(cartServ)

	cartRoutes := router.Group("cart")

	cartRoutes.POST("/add-cart", cartHand.CreateShopCart)
	cartRoutes.GET("/get/:id", cartHand.GetCartByID)
}
