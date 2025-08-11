package service

import (
	ctgP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/port"
	productsP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/port"
)

var _ ctgP.CategorySrv = &Service{}

type Service struct {
	CategoryRepo ctgP.CategoryRepo
	ProductRepo productsP.ProductRepo
}

func NewCategorySrv(categoryRepo ctgP.CategoryRepo, productRepo productsP.ProductRepo) *Service {
	return &Service{
		CategoryRepo: categoryRepo,
		ProductRepo: productRepo,
	}
}
