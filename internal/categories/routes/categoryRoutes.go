package routes

import (
	"e-commerce/internal/categories/handler"
	"e-commerce/internal/categories/repository"
	"e-commerce/internal/categories/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitCategoryRoutes(router *gin.RouterGroup, db *sqlx.DB) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/add", categoryHandler.CreateCategory)
	categoryRoutes.DELETE("/delete/:id", categoryHandler.DeleteCategoryByID)
	categoryRoutes.PUT("/update/:id", categoryHandler.UpdateCategoryByID)
}
