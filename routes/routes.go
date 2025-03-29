package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nbursa/user-admin-api/controllers"
	"github.com/nbursa/user-admin-api/repositories"
	"github.com/nbursa/user-admin-api/services"
)

func RegisterRoutes(r *gin.Engine) {
	repo := repositories.NewUserMongoRepository()
	service := services.NewUserService(repo)
	controller := controllers.NewUserController(service)

	r.POST("/users", controller.CreateUser)
}
