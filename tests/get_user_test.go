package user_service

import (
	"context"
	"errors"
	"testing"

	"github.com/nbursa/user-admin-api/models"
	"github.com/nbursa/user-admin-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserByID_Success(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	expectedUser := &models.User{
		ID:    [12]byte{1, 2, 3},
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	}

	repo.On("GetByID", mock.Anything, "valid-id").Return(expectedUser, nil)

	user, err := service.GetUserByID(context.Background(), "valid-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	repo.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	repo.On("GetByID", mock.Anything, "missing-id").Return((*models.User)(nil), errors.New("not found"))

	user, err := service.GetUserByID(context.Background(), "missing-id")

	assert.Nil(t, user)
	assert.Error(t, err)
	repo.AssertExpectations(t)
}