package user_service

import (
	"context"
	"errors"
	"testing"

	"github.com/nbursa/user-admin-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUserByID_Success(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	repo.On("DeleteByID", mock.Anything, "user123").Return(nil)

	err := service.DeleteUserByID(context.Background(), "user123")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDeleteUserByID_NotFound(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	repo.On("DeleteByID", mock.Anything, "nonexistent").Return(errors.New("mongo: no documents in result"))

	err := service.DeleteUserByID(context.Background(), "nonexistent")

	assert.Error(t, err)
	assert.Equal(t, "mongo: no documents in result", err.Error())
	repo.AssertExpectations(t)
}