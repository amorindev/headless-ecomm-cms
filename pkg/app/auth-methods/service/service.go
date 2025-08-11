package service

import (
	authMethodP "github.com/amorindev/headless-ecomm-cms/pkg/app/auth-methods/port"
	otpCodeP "github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/port"
	phoneP "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/port"
	roleP "github.com/amorindev/headless-ecomm-cms/pkg/app/roles/port"
	sessionP "github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/port"
	userP "github.com/amorindev/headless-ecomm-cms/pkg/app/users/port"
	emailP "github.com/amorindev/headless-ecomm-cms/pkg/email/port"
	smsP "github.com/amorindev/headless-ecomm-cms/pkg/sms/port"
)

var _ authMethodP.AuthMethodSrv = &Service{}

type Service struct {
	RoleRepo    roleP.RoleRepo
	UserRepo    userP.UserRepo
	OtpRepo     otpCodeP.OtpRepo
	SessionRepo sessionP.SessionRepo
	PhoneRepo   phoneP.PhoneRepo
	EmailSrv    emailP.EmailSrv
	SmsSrv      smsP.SmsSrv
	SessionSrv  sessionP.SessionSrv
}

func NewAuthMethodSrv(
	userRepo userP.UserRepo,
	roleRepo roleP.RoleRepo,
	otpRepo otpCodeP.OtpRepo,
	sessionRepo sessionP.SessionRepo,
	phoneRepo phoneP.PhoneRepo,
	sessionSrv sessionP.SessionSrv,
	emailSrv emailP.EmailSrv,
	smsSrv smsP.SmsSrv,
) *Service {
	return &Service{
		RoleRepo:    roleRepo,
		UserRepo:    userRepo,
		OtpRepo:     otpRepo,
		SessionRepo: sessionRepo,
		PhoneRepo:   phoneRepo,
		EmailSrv:    emailSrv,
		SmsSrv:      smsSrv,
		SessionSrv:  sessionSrv,
	}
}
