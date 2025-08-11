package service

import (
	"context"
	"math"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/helpers"
)

func (s *Service) GetAll(ctx context.Context, limit int64, page int64) ([]*domain.Product, int64, int64, error) {

	products, err := s.ProductRepo.FindAll(ctx, limit, page)
	if err != nil {
		return nil, 0, 0, err
	}

	for _, product := range products {
		product.CategoryID = nil

		url, err := s.FileStorageSrv.GetImage(ctx, product.FilePath)
		if err != nil {
			return nil, 0, 0, err
		}

		product.ImgUrl = url

		for _, pItem := range product.ProductItems {
			productURL, err := s.FileStorageSrv.GetImage(ctx, pItem.FilePath)
			if err != nil {
				return nil, 0, 0, err
			}
			pItem.ImgUrl = productURL
			pItem.FilePath = ""
			pItem.VarOptionIDs = nil
		}
		product.FilePath = ""
		helpers.CalculateVariations(product)
	}

	count, err := s.ProductRepo.Count(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	totalPages := int64(math.Ceil(float64(count) / float64(limit)))

	return products, count, totalPages, nil
}
