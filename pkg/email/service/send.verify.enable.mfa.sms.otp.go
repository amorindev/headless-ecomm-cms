package service

import "github.com/amorindev/headless-ecomm-cms/pkg/email/service/templates"

type VerifyEnableMfaOtpTmplData struct {
	Name    string
	Subject string
	Code    string
}

func (s *EmailService) SendVerifyEnableMfaSmsOtp(email string, code string) error {
	
	data := VerifyEnableMfaOtpTmplData{
		Name:    "MyCompany",
		Subject: email,
		Code:    code,
	}

	tmplString, err := templates.LoadTemplate("pkg/email/service/templates/verify-enable-mfa-sms-otp.html", data)
	if err != nil {
		return err
	}
	
	subject :=  "Verify Your Phone Number to Enable SMS MFA"
	return s.EmailAdapter.Send(email, subject,tmplString)
}
