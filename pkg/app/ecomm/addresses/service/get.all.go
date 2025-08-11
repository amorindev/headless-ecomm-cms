package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
)

func (s *Service) GetAll(ctx context.Context, userID string) ([]*domain.Address, error) {
	return s.AddressRepo.GetAll(ctx, userID)
}
