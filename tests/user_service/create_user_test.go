package user_service

import (
	"context"
	"testing"

	"github.com/nbursa/user-admin-api/models"
	"github.com/nbursa/user-admin-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	userInput := models.UserInput{
		Name:  "John Doe",
		Email: "john@doe.com",
		Age:   30,
	}

	repo.On("ExistsByEmail", mock.Anything, userInput.Email).Return(false, nil)
	repo.On("Insert", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	user, err := service.CreateUser(context.Background(), userInput)

	assert.NoError(t, err)
	assert.Equal(t, userInput.Email, user.Email)
	assert.Equal(t, userInput.Name, user.Name)
	assert.Equal(t, userInput.Age, user.Age)
	repo.AssertExpectations(t)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	userInput := models.UserInput{
		Name:  "John Duplicate",
		Email: "duplicate@doe.com",
		Age:   28,
	}

	repo.On("ExistsByEmail", mock.Anything, userInput.Email).Return(true, nil)

	user, err := service.CreateUser(context.Background(), userInput)

	assert.Nil(t, user)
	assert.Error(t, err)
	assert.Equal(t, "email already in use", err.Error())
	repo.AssertExpectations(t)
}
