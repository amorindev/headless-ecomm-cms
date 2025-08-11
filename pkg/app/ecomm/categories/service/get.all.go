package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
)

func (s *Service) GetAll(ctx context.Context) ([]*domain.Category, error) {
	categories, err := s.CategoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
