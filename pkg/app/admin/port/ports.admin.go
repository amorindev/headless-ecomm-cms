package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
)

type AdminRepo interface {
	Exists(ctx context.Context, roleAdminID string) (bool, error)
}

type AdminSrv interface{
    Create(ctx context.Context, user *domain.User) error
}
