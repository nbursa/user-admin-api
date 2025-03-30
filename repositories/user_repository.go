package repositories

import (
	"context"

	"github.com/nbursa/user-admin-api/models"
)

type UserRepository interface {
	Insert(ctx context.Context, user *models.User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByNameOrEmailPaginated(ctx context.Context, search string, page, limit int) ([]*models.User, int, error)
}
