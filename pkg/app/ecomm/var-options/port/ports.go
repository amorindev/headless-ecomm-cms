package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
)

type VariationOptionRepo interface {
	Insert(ctx context.Context, varOption *domain.VariationOption) error
	FindByVariationID(ctx context.Context, variationID string) ([]*domain.VariationOption, error)
	Delete(ctx context.Context, id string) error
	FindByName(ctx context.Context, name string) (*domain.VariationOption, error)
	Exists(ctx context.Context, name string) (bool, error)
}
