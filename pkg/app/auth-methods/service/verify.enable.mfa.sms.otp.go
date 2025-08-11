package service

import (
	"context"
	"errors"
	"time"

	authCts "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/constants"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"

)

func (s *Service) VerifyEnableMfaSmsOtp(ctx context.Context, otpID string, otpCode string, userID string) error {
	user, err := s.UserRepo.Find(ctx, userID)
	if err != nil {
		return err
	}

	otp, err := s.OtpRepo.Find(ctx, otpID)
	if err != nil {
		return err
	}

	if time.Now().After(*otp.ExpiresAt) {
		return errors.New("otp-expired")
	}

	if otpCode != otp.OptCode {
		return errors.New("otp-code-do-not-match")
	}

	if otp.Purpose != authCts.VerifyEnableMfaSmsPurpose {
		return errors.New("otp-invalid-purpose")
	}
	err = s.UserRepo.ConfirmMfaSms(ctx, userID)
	if err != nil {
		return err
	}

	if user.MfaStatus == nil {
		user.MfaStatus = &domain.MfaStatus{
			IsMfaSmsCompleted: false,
		}

		err := s.UserRepo.UpdateMfaStatus(ctx, userID, user.MfaStatus)
		if err != nil {
			return err
		}
	}

	err = s.OtpRepo.Delete(ctx, otp.ID.(string))
	if err != nil {
		return err
	}

	return nil
}
