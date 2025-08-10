package service

import "github.com/amorindev/headless-ecomm-cms/pkg/email/port"

var _ port.EmailSrv = &EmailService{}


type EmailService struct {
	EmailAdapter port.EmailAdapter
}

func NewEmailSrv(emailAdapter port.EmailAdapter) *EmailService {
    return &EmailService{
		EmailAdapter: emailAdapter,
	}
}
