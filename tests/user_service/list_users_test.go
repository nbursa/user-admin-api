package user_service

import (
	"context"
	"testing"

	"github.com/nbursa/user-admin-api/models"
	"github.com/nbursa/user-admin-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsersPaginated(t *testing.T) {
	repo := new(mockUserRepo)
	service := services.NewUserService(repo)

	expectedUsers := []*models.User{
		{Name: "Tester", Email: "tester@example.com", Age: 30},
		{Name: "Tester2", Email: "tester2@example.com", Age: 35},
	}

	repo.On("FindByNameOrEmailPaginated", mock.Anything, "a", 1, 2).Return(expectedUsers, 10, nil)

	result, err := service.GetUsersPaginated(context.Background(), "a", 1, 2)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Users))
	assert.Equal(t, 10, result.Total)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 2, result.Limit)

	repo.AssertExpectations(t)
}
