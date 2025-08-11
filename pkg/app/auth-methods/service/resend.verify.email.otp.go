package service

import (
	"context"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/constants"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/utils"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/domain"
)

func (s *Service) ResendVerifyEmailOtp(ctx context.Context, email string) (string, error) {
	user, err := s.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	
	code, err := utils.GenOtpCode()
	if err != nil {
		return "", err
	}

	now := time.Now()
	purpose := constants.VerifyEmailPurpose
	expiresAt := now.Add(time.Hour)
	used := false

	otp := &domain.OtpCodes{
		UserID:    user.ID.(string),
		OptCode:   code,
		Purpose:   purpose,
		Used:      used,
		ExpiresAt: &expiresAt,
		CreatedAt: &now,
	}

	err = s.OtpRepo.Insert(context.Background(), otp)
	if err != nil {
		return "", err
	}

	err = s.EmailSrv.SendVerifyEmailOtp(email, code)
	if err != nil {
		return "", err
	}

	return otp.ID.(string), nil
}
