package service

import "github.com/amorindev/headless-ecomm-cms/pkg/email/service/templates"

type VerifyMfaSmsOtpTmplData struct {
	Name    string
	Subject string
	Code    string
}

func (s *EmailService) SendVerifyMfaSmsOtp(email string, code string) error {

	data := VerifyMfaSmsOtpTmplData{
		Name:    "MyCompany",
		Subject: email,
		Code:    code,
	}

	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/verify-mfa-sms-otp.html", data)
	if err != nil {
		return err
	}

	subject := "Your SMS Verification Code for Sign-In"

	return s.EmailAdapter.Send(email, subject, tmplString)
}
