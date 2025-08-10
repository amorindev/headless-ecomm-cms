package resend

import (
	"github.com/resend/resend-go/v2"
)

func (a *ResendAdapter) Send(to, subject, htmlBody string) error {
	params := &resend.SendEmailRequest{
		From:    a.Form,
		To:      []string{to},
		Subject: subject,
		Html:    htmlBody,
	}

	_, err := a.Client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}
