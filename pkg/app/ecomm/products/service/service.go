package service

import (

	fileStgP "github.com/amorindev/headless-ecomm-cms/pkg/file-storage/port"
	productP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/port"
	categoryP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/port"
)

var _ productP.ProductSrv = &Service{}

type Service struct {
	ProductRepo    productP.ProductRepo
	CategoryRepo   categoryP.CategoryRepo
	FileStorageSrv fileStgP.FileStorageSrv
}

func NewProductSrv(
	productRepo productP.ProductRepo,
	categoryRepo categoryP.CategoryRepo,
	fileStorageSrv fileStgP.FileStorageSrv,
) *Service {
	return &Service{
		ProductRepo:    productRepo,
		CategoryRepo:   categoryRepo,
		FileStorageSrv: fileStorageSrv,
	}
}
