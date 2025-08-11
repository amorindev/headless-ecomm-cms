package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
)

type ProductItemRepo interface {
	Insert(ctx context.Context, productItem *domain.ProductItem) error
}
