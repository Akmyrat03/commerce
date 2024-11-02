package main

import (
	userRoutes "e-commerce/internal/users/routes"
	"e-commerce/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
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

	app := gin.Default()
	api := app.Group("/api")
	userRoutes.InitUserRoutes(api, DB)

	if err := app.Run("localhost:8000"); err != nil {
		log.Fatalf("Failed running app: %v", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
