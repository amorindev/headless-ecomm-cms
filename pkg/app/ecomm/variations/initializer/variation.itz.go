package initializer

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/port"
)

type VariationItz struct {
	VariationRepo port.VariationRepo
}

func NewVariationItz(variationRepo port.VariationRepo) *VariationItz {
	return &VariationItz{
		VariationRepo: variationRepo,
	}
}

func (vi *VariationItz) SeedEssential(ctx context.Context) error{
    variations := []*domain.Variation{
        domain.NewVariation("Color"),
        domain.NewVariation("Size"),
    }

    for _, v := range variations {
        exists, err := vi.VariationRepo.Exists(ctx, v.Name)
        if err != nil {
			return err
		}
        if !exists {
			if err := vi.VariationRepo.Insert(ctx, v); err != nil {
				return err
			}
		}
    }
    return nil
}

