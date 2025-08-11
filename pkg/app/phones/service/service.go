package service

import "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/port"

var _ port.PhoneSrv = &Service{}

type Service struct {
	PhoneRepo port.PhoneRepo
}

func NewPhoneService(phoneRepo port.PhoneRepo) *Service {
	return &Service{
		PhoneRepo: phoneRepo,
	}
}
