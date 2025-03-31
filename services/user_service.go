package services

import (
	"context"
	"errors"

	"github.com/nbursa/user-admin-api/models"
	"github.com/nbursa/user-admin-api/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, input models.UserInput) (*models.User, error)
	GetUsersPaginated(ctx context.Context, search string, page, limit int) (*models.PaginatedUsers, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	DeleteUserByID(ctx context.Context, id string) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	exists, err := s.repo.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("email already in use")
	}

	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}

	if err := s.repo.Insert(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUsersPaginated(ctx context.Context, search string, page, limit int) (*models.PaginatedUsers, error) {
	users, total, err := s.repo.FindByNameOrEmailPaginated(ctx, search, page, limit)
	if err != nil {
		return nil, err
	}

	return &models.PaginatedUsers{
		Users: users,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) DeleteUserByID(ctx context.Context, id string) error {
	return s.repo.DeleteByID(ctx, id)
}
