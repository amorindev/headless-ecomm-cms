package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
)

func (s *Service) Create(ctx context.Context, variation *domain.Variation) error {
    return s.VariationRepo.Insert(ctx, variation)
}
