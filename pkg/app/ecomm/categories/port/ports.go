package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
)

type CategoryRepo interface {
	Insert(ctx context.Context, category *domain.Category) error
	FindByName(ctx context.Context, name string) (*domain.Category, error)
	FindAll(ctx context.Context) ([]*domain.Category, error)
	Exists(ctx context.Context, name string) (bool, error)
	Update(ctx context.Context, id string, name string) error
	Delete(ctx context.Context, id string) error
}

type CategorySrv interface {
	Create(ctx context.Context, category *domain.Category) error
	GetAll(ctx context.Context) ([]*domain.Category, error) 
	Update(ctx context.Context, id string, name string) error
	Delete(ctx context.Context, id string) error
}
