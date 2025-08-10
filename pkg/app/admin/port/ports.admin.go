package port

import (
	"context"

)

type AdminRepo interface {
	Exists(ctx context.Context, roleAdminID string) (bool, error)
}

type AdminSrv interface{
    Create(ctx context.Context, user *domain.User) error
}
