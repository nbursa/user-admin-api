package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nbursa/user-admin-api/config"
	"github.com/nbursa/user-admin-api/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting server...")

	// Load .env
	if err := godotenv.Load(); err != nil {
		logrus.Warn(".env file not found or failed to load")
	}

	config.ConnectMongo()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_ORIGIN")},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
