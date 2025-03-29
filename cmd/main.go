package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nbursa/user-admin-api/config"
	"github.com/nbursa/user-admin-api/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting server...")

	config.ConnectMongo()
	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
