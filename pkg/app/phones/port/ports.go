package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
)

type PhoneRepo interface {
	Insert(ctx context.Context, phone *domain.Phone) error
	Find(ctx context.Context, id string) (*domain.Phone, error)
	FindAllByUserID(ctx context.Context, userID string) ([]*domain.Phone, error)
	FindDefault(ctx context.Context) (*domain.Phone, error)
}

type PhoneSrv interface {
	Create(ctx context.Context, phone *domain.Phone) error
	GetAllByUserID(ctx context.Context, userID string) ([]*domain.Phone, error)
}
