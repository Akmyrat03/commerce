package main

import (
	"context"
	cartRoutes "e-commerce/internal/cart/routes"
	cartItemRoutes "e-commerce/internal/cart_items/routes"
	categoryRoutes "e-commerce/internal/categories/routes"
	prodRoutes "e-commerce/internal/products/routes"
	userRoutes "e-commerce/internal/users/routes"
	"e-commerce/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	DB, err := database.ConnectToDB(database.Config{
		Host:     viper.GetString("DB.host"),
		Port:     viper.GetString("DB.port"),
		Username: viper.GetString("DB.username"),
		Password: viper.GetString("DB.password"),
		DBName:   viper.GetString("DB.dbname"),
		SSLMode:  viper.GetString("DB.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err.Error())
	}

	redisAddr := viper.GetString("REDIS.addr")
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ctx := context.Background()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	app := gin.Default()
	api := app.Group("/api")

	userRoutes.InitUserRoutes(api, DB, redisClient)
	categoryRoutes.InitCategoryRoutes(api, DB)
	prodRoutes.InitRoutes(DB, api)
	cartRoutes.InitCartRoutes(DB, api)
	cartItemRoutes.InitCartRoutes(DB, api)

	if err := app.Run(viper.GetString("APP.host")); err != nil {
		log.Fatalf("Failed running app: %v", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("REDIS.addr", "localhost:6379")
	viper.SetDefault("APP.host", "localhost:8000")
	return nil
}
