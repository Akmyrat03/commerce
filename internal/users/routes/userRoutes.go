package routes

import (
	"e-commerce/internal/users/handler"
	"e-commerce/internal/users/repository"
	"e-commerce/internal/users/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitUserRoutes(router *gin.RouterGroup, DB *sqlx.DB) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserService(userService)

	userRoutes := router.Group("/users")
	userRoutes.POST("/sign-up", userHandler.SignUp)
	userRoutes.POST("/sign-in", userHandler.Login)
	userRoutes.DELETE("/sign-out", userHandler.SignOut)
}
