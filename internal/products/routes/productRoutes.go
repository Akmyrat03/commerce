package routes

import (
	"e-commerce/internal/products/handler"
	"e-commerce/internal/products/repository"
	"e-commerce/internal/products/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(db *sqlx.DB, router *gin.RouterGroup) {
	prodRepo := repository.NewProductRepository(db)
	prodService := service.NewProductService(prodRepo)
	prodHandler := handler.NewProductHandler(prodService)

	prodRoutes := router.Group("/products")

	prodRoutes.POST("/add-product", prodHandler.CreateProduct)
	prodRoutes.DELETE("/delete/:id", prodHandler.DeleteProduct)
	prodRoutes.GET("/view-all", prodHandler.GetAllProducts)
	prodRoutes.GET("/view/:name", prodHandler.GetProductByCategoryName)
}
