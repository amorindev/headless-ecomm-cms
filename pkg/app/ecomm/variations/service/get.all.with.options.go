package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
)

func (s *Service) GetAllWithOptions(ctx context.Context) ([]*domain.Variation, error) {
	return s.VariationRepo.FindAllWithOptions(ctx)
}
