package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
)

func (s *Service) CreateOption(ctx context.Context, varOption *domain.VariationOption) error {
	return s.VarOptRepo.Insert(ctx, varOption)
}
