package service

import (
	"context"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
)

func (s *Service) Create(ctx context.Context, address *domain.Address) error {
	now := time.Now().UTC()
	address.CreatedAt = &now
	address.UpdatedAt = &now

	address.IsDefault = false

	return s.AddressRepo.Insert(ctx, address)
}
