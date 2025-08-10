package service

import "github.com/amorindev/headless-ecomm-cms/pkg/email/service/templates"

type VerifyEmailOtpTmplData struct {
	Name    string
	Subject string
	Code    string
}

func (s *EmailService) SendVerifyEmailOtp(email string, code string) error {

	data := VerifyEmailOtpTmplData{
		Name:    "MyCompany",
		Subject: email,
		Code:    code,
	}

	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/verify-email-otp.html", data)
	if err != nil {
		return err
	}

	subject := "Verify Your Email Address"

	return s.EmailAdapter.Send(email, subject,tmplString)
}
