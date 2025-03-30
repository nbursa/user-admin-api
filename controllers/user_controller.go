package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nbursa/user-admin-api/models"
	"github.com/nbursa/user-admin-api/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.CreateUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	search := c.Query("search")
	page := parseQueryInt(c, "page", 1)
	limit := parseQueryInt(c, "limit", 10)

	result, err := uc.service.GetUsersPaginated(c.Request.Context(), search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func parseQueryInt(c *gin.Context, key string, defaultVal int) int {
	valStr := c.Query(key)
	if val, err := strconv.Atoi(valStr); err == nil && val > 0 {
		return val
	}
	return defaultVal
}
