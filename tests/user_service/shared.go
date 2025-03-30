package user_service

import (
	"context"

	"github.com/nbursa/user-admin-api/models"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) Insert(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *mockUserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	return args.Bool(0), args.Error(1)
}

func (m *mockUserRepo) FindByNameOrEmailPaginated(ctx context.Context, search string, page, limit int) ([]*models.User, int, error) {
	args := m.Called(ctx, search, page, limit)
	return args.Get(0).([]*models.User), args.Int(1), args.Error(2)
}
