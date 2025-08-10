package service

import "github.com/amorindev/headless-ecomm-cms/pkg/sms/port"

var _ port.SmsSrv = &Service{}

type Service struct {
	SmsAdp port.SmsAdp
}

func NewSmsSrv(smsAdp port.SmsAdp) *Service {
	return &Service{
		SmsAdp: smsAdp,
	}
}
