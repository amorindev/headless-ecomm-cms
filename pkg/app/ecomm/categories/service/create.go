package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
)

func (s *Service) Create(ctx context.Context, category *domain.Category) error {
	return s.CategoryRepo.Insert(ctx, category)
}
