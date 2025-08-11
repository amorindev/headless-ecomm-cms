package core

import (
	"errors"
)

type PhoneIDReq struct {
	PhoneID string `json:"phone_id"`
}

func (req *PhoneIDReq) IsPhoneIDValid() error {
	if req.PhoneID == "" {
		return errors.New("phone_id field is required")
	}
	return nil
}

type UserIDReq struct {
	UserID string `json:"user_id"`
}

func (req *UserIDReq) IsUserIDValid() error {
	if req.UserID == "" {
		return errors.New("user_id field is required")
	}
	return nil
}

type OtpIDAndCodeReq struct {
	OtpID   string `json:"otp_id"`
	OtpCode string `json:"otp_code"`
}

func (req *OtpIDAndCodeReq) IsOtpIDAndCodeValid() error{
	if req.OtpID == "" {
		return errors.New("otp_id field is required")
	}
	if req.OtpCode == "" {
		return errors.New("otp_code field is required")
	}
	return nil
}
