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

func TestUpdateUserByID_Success(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	update := &models.UserInput{
		Name:  "Updated Name",
		Email: "updated@example.com",
		Age:   35,
	}

	repo.On("ExistsByEmail", mock.Anything, update.Email).Return(false, nil)
	repo.On("UpdateByID", mock.Anything, "user-id", update).Return(nil)

	err := service.UpdateUserByID(context.Background(), "user-id", update)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestUpdateUserByID_EmailExistsOtherUser(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	update := &models.UserInput{
		Name:  "Updated Name",
		Email: "existing@example.com",
		Age:   40,
	}

	repo.On("ExistsByEmail", mock.Anything, update.Email).Return(true, nil)
	repo.On("GetByID", mock.Anything, "user-id").Return(&models.User{
		ID:    [12]byte{1},
		Name:  "Old Name",
		Email: "old@example.com",
		Age:   30,
	}, nil)

	err := service.UpdateUserByID(context.Background(), "user-id", update)

	assert.Error(t, err)
	assert.Equal(t, "email already in use", err.Error())
	repo.AssertExpectations(t)
}

func TestUpdateUserByID_NotFound(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	update := &models.UserInput{
		Name:  "Name",
		Email: "email@example.com",
		Age:   20,
	}

	repo.On("ExistsByEmail", mock.Anything, update.Email).Return(false, nil)
	repo.On("UpdateByID", mock.Anything, "missing-id", update).Return(errors.New("not found"))

	err := service.UpdateUserByID(context.Background(), "missing-id", update)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
	repo.AssertExpectations(t)
}