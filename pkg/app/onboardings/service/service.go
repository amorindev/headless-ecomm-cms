package service

import (
    obdP "github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/port"
    fileStgP "github.com/amorindev/headless-ecomm-cms/pkg/file-storage/port"
)

var _ obdP.OnboardingSrv = &Service{}

type Service struct {
	OnboardingRepo obdP.OnboardingRepo
	FileStorageSrv fileStgP.FileStorageSrv
}

func NewOnboardingSrv(onboardingRepo obdP.OnboardingRepo, fileStorageSrv fileStgP.FileStorageSrv) *Service {
	return &Service{
		OnboardingRepo: onboardingRepo,
		FileStorageSrv: fileStorageSrv,
	}
}
