package service

import (
	varOptPort "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/port"
	variationPort "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/port"
	productPort "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/port"
	
)

var _ variationPort.VariationSrv = &Service{}


type Service struct {
	VariationRepo   variationPort.VariationRepo
	VarOptRepo      varOptPort.VariationOptionRepo
	ProductItemRepo productPort.ProductItemRepo
}

func NewVariationSrv(variationRepo variationPort.VariationRepo,varOptRepo varOptPort.VariationOptionRepo) *Service {
	return &Service{
		VariationRepo: variationRepo,
		VarOptRepo: varOptRepo,
	}
}
