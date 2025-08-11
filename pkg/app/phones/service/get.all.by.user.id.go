package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
)

func (s *Service) GetAllByUserID(ctx context.Context, userID string) ([]*domain.Phone, error) {
	return s.PhoneRepo.FindAllByUserID(ctx, userID)
}
