package service

import "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/port"

var _ port.AddressSrv = &Service{}

type Service struct {
	AddressRepo port.AddressRepo
}

func NewAddressSrv(addressRepo port.AddressRepo) *Service {
	return &Service{
		AddressRepo: addressRepo,
	}
}
