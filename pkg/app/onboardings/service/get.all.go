package service

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
)

func (s *Service) GetAll(ctx context.Context) ([]*domain.Onboarding, error) {
	onboardings, err := s.OnboardingRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, onboarding := range onboardings {
		url, err := s.FileStorageSrv.GetImage(ctx, onboarding.FilePath)
		if err != nil {
			return nil, err
		}
		onboarding.ImgUrl = url
		onboarding.FilePath = ""
	}

	return onboardings, nil
}
