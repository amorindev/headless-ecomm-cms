package service

import (
	"context"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/constants"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/utils"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/domain"
)

func (s *Service) ResendVerifyMfaSmsOtp(ctx context.Context, userID string) (string, error) {
	// * verify if user exists
	user, err := s.UserRepo.Find(ctx, userID)
	if err != nil {
		return "", err
	}

	code, err := utils.GenOtpCode()
	if err != nil {
		return "", err
	}

	purpose := constants.VerifyMfaSmsPurpose
	now := time.Now()
	expiresAt := now.Add(time.Hour)
	used := false

	otp := &domain.OtpCodes{
		OptCode:   code,
		Purpose:   purpose,
		Used:      used,
		UserID:    user.ID,
		ExpiresAt: &expiresAt,
		CreatedAt: &now,
	}

	err = s.OtpRepo.Insert(ctx, otp)
	if err != nil {
		return "", err
	}

	mfaSms, err := s.UserRepo.FindMfaSmsByUserID(ctx, user.ID.(string))
	if err != nil {
		return "", err
	}

	phone, err := s.PhoneRepo.Find(ctx, mfaSms.PhoneID.(string))
	if err != nil {
		return "", err
	}

	err = s.SmsSrv.SendMfaSms(phone.CountryCode, phone.Number, otp.OptCode)
	if err != nil {
		return "", err
	}

	return otp.ID.(string), nil
}
