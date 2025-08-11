package service

import (
	"context"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
)

func (s *Service) Create(ctx context.Context, phone *domain.Phone) error {
	now := time.Now().UTC()

	phone.CreatedAt = &now
	phone.UpdatedAt = &now
	phone.IsVerified = false
	phone.IsDefault = false

	return s.PhoneRepo.Insert(ctx, phone)
}
