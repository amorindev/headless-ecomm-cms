package service

import (
	"context"
	"errors"

	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/constants"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/service/utils"
	otpCodeD "github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/domain"
	userD "github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	userErr "github.com/amorindev/headless-ecomm-cms/pkg/app/users/errors"
)

func (s *Service) EnableMfaSms(ctx context.Context, userID string, phoneID string) (string, error) {
	mfaSms, err := s.UserRepo.FindMfaSmsByUserID(ctx, userID)
	if err != nil {
		if err != userErr.ErrMfaSmsNotFound{
			return "", err
		}
	}

	if mfaSms != nil && mfaSms.Confirmed {
		return "", errors.New("MFA via SMS is already enabled")
	} 

	phone, err := s.PhoneRepo.Find(ctx, phoneID)
	if err != nil {
		return "", err
	}

	newMfaSms := userD.NewUserMfaSms(userID, phoneID, false)

	err = s.UserRepo.InsertMfaSms(ctx, newMfaSms)
	if err != nil {
		return "", err
	}

	code, err := utils.GenOtpCode()
	if err != nil {
		return "", err
	}

	purpose := constants.VerifyEnableMfaSmsPurpose
	now := time.Now()
	expiresAt := now.Add(time.Hour)
	used := false
	otp := &otpCodeD.OtpCodes{
		OptCode:   code,
		Purpose:   purpose,
		ExpiresAt: &expiresAt,
		Used:      used,
		CreatedAt: &now,
	}

	otp.UserID = userID

	err = s.OtpRepo.Insert(ctx, otp)
	if err != nil {
		return "", err
	}

	err = s.SmsSrv.SendEnableMfaSms(phone.CountryCode, phone.Number, otp.OptCode)
	if err != nil {
		return "", err
	}

	return otp.ID.(string), nil
}
