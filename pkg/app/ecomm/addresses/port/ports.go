package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
)

type AddressRepo interface {
	Insert(ctx context.Context, address *domain.Address) error
	Get(ctx context.Context, id string) (*domain.Address, error)
	GetAll(ctx context.Context, userID string) ([]*domain.Address, error)
	GetDefault(ctx context.Context) (*domain.Address, error)
}

type AddressSrv interface {
	Create(ctx context.Context, address *domain.Address) error
	GetAll(ctx context.Context, userID string) ([]*domain.Address, error)
}

