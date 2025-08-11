package service

import (
	"context"
	"fmt"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/helpers"
)

func (s *Service) Create(ctx context.Context, onboarding *domain.Onboarding) error {
	now := time.Now().UTC()

	onboarding.CreatedAt = &now
	onboarding.Seen = false

	// * Create a unique file name to avoid overwriting existing files
	uniqueFileName := helpers.GenerateUniqueFileName(onboarding.FilePath)

	onboarding.FilePath = uniqueFileName

	bucketFolderStruct := "onboardings/"
	onboarding.FilePath = fmt.Sprintf("%s%s", bucketFolderStruct, onboarding.FilePath)

	err := s.FileStorageSrv.UploadImage(context.Background(), onboarding.FilePath, onboarding.File, onboarding.ContentType)
	if err != nil {
		return err
	}

	err = s.OnboardingRepo.Insert(ctx, onboarding)
	if err != nil {
		return err
	}

	return nil
}
