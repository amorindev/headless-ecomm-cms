package port

import (
	"context"

	varOptD "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
)

type VariationRepo interface {
	Insert(ctx context.Context, variation *domain.Variation) error
	FindAll(ctx context.Context) ([]*domain.Variation, error)
	FindAllWithOptions(ctx context.Context) ([]*domain.Variation, error)
	Delete(ctx context.Context, id string) error
	FindByName(ctx context.Context, name string) (*domain.Variation, error)
	Exists(ctx context.Context, name string) (bool, error)
}

type VariationSrv interface {
	// * Variations
	Create(ctx context.Context, variation *domain.Variation) error
	GetAllWithOptions(ctx context.Context) ([]*domain.Variation, error)
	Delete(ctx context.Context, id string) error

	// * variation options
	CreateOption(ctx context.Context, varOption *varOptD.VariationOption) error
	DeleteOption(ctx context.Context, varOptID string) error
}
