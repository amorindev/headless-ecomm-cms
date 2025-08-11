package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
)

type ProductRepo interface {
	Find(ctx context.Context, id string) (*domain.Product, error)

	Insert(ctx context.Context, product *domain.Product) error

	FindByName(ctx context.Context, name string) (*domain.Product, error)

	FindAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, error)

	Exists(ctx context.Context, name string) (bool, error)

	Count(ctx context.Context) (int64, error)

	FindByCategory(ctx context.Context, categoryID string) ([]*domain.Product, error)
}

type ProductSrv interface {
	GetAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, int64, int64, error)

	CreateFromZip(ctx context.Context, product *domain.Product) error
}
