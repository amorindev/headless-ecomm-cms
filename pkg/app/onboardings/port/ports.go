package port

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
)

type OnboardingRepo interface {
	Insert(ctx context.Context, onboarding *domain.Onboarding) error
	GetAll(ctx context.Context) ([]*domain.Onboarding, error)
	GetByTitle(ctx context.Context, title string) (*domain.Onboarding, error)
}

type OnboardingSrv interface {
	Create(ctx context.Context, onboarding *domain.Onboarding) error
	GetAll(ctx context.Context) ([]*domain.Onboarding, error)
}


