package routes

import (
	"e-commerce/internal/users/middleware"
	"e-commerce/internal/users/repository"
	"e-commerce/internal/users/service"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func InitUserRoutes(router *gin.RouterGroup, DB *sqlx.DB, redisClient *redis.Client) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userMiddleware := middleware.NewUserMiddleware(userRepo, userService, redisClient)

	userRoutes := router.Group("/users")

	userRoutes.POST("/sign-up", userMiddleware.SignUp())
	userRoutes.POST("/login", userMiddleware.Login())
	userRoutes.DELETE("/sign-out", userMiddleware.SignOut())
	userRoutes.GET("/profile", userMiddleware.Authenticate(), userMiddleware.Profile())
	userRoutes.GET("/view-users", userMiddleware.GetAllUsers())

}
