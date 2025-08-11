package core

import "errors"

type VerifyEmailOTPReq struct {
	OtpID   string `json:"otp_id"`
	OtpCode string `json:"otp_code"`
	UserID  string `json:"user_id"`
}

func (req *VerifyEmailOTPReq) IsVerifyEmailOTPValid() error {
	if req.OtpID == "" {
		return errors.New("otp_id field is required")
	}
	if req.OtpCode == "" {
		return errors.New("otp_code field is required")
	}

	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	return nil
}
