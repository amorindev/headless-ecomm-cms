package service

import (
	"context"
	"errors"
	"time"


	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/constants"
	sessionD "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"

)

func (s *Service) VerifyMfaSmsOtp(ctx context.Context, otpID string, otpCode string) (*userD.User, *sessionD.Session, error) {
	otp, err := s.OtpRepo.Find(ctx, otpID)
	if err != nil {
		return nil, nil, err
	}

	if time.Now().After(*otp.ExpiresAt) {
		return nil, nil, errors.New("otp-expired")
	}

	if otpCode != otp.OptCode {
		return nil, nil, errors.New("otp-code-do-not-match")
	}

	if otp.Purpose != constants.VerifyMfaSmsPurpose {
		return nil, nil, errors.New("otp-invalid-purpose")
	}

	user, err := s.UserRepo.Find(ctx, otp.UserID.(string))
	if err != nil {
		return nil, nil, err
	}

	user.MfaStatus.ResetMfaStatus()
	err = s.UserRepo.ResetMfaStatus(ctx, user.ID.(string), user.MfaStatus)
	if err != nil {
		return nil, nil, err
	}

	// *  Create the session
	
	// * Assign roles
	roles, err := s.RoleRepo.FindByIDs(ctx, user.RolesIDs)
	if err != nil {
		return nil, nil, err
	}

	user.Roles = roles

	session := &sessionD.Session{
		UserID:     user.ID.(string),
		RememberMe: false,
	}

	err = s.SessionSrv.Create(session, roles, user.Email)
	if err != nil {
		return nil, nil, err
	}

	err = s.OtpRepo.Delete(ctx, otp.ID.(string))
	if err != nil {
		return nil, nil, err
	}

	return user, session, nil
}
