package service

import (
	"context"
	"fmt"
	"time"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/helpers"
)

func (s *Service) CreateFromZip(ctx context.Context, product *domain.Product) error {
	exists, err := s.ProductRepo.Exists(ctx, product.Name)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	ctg, err := s.CategoryRepo.FindByName(context.Background(), product.CategoryName)
	if err != nil {
		return err
	}

	product.CategoryID = ctg.ID

	now := time.Now().UTC()
	product.CreatedAt = &now
	product.UpdatedAt = &now

	// * Create a unique file name to avoid overwriting existing files
	uniqueFileName := helpers.GenerateUniqueFileName(product.FilePath)

	product.FilePath = uniqueFileName

	bucketFolderStruct := "products/"
	product.FilePath = fmt.Sprintf("%s%s", bucketFolderStruct, product.FilePath)

	err = s.FileStorageSrv.UploadImage(ctx, product.FilePath, product.File, product.ContentType)
	if err != nil {
		return err
	}

	for i, p := range product.ProductItems {

		product.ProductItems[i].CreatedAt = &now
		product.ProductItems[i].UpdatedAt = &now

		uniqueName := helpers.GenerateUniqueFileName(p.FilePath)
		product.ProductItems[i].FilePath = uniqueName

		product.ProductItems[i].FilePath = bucketFolderStruct + product.ProductItems[i].FilePath

		product.ProductItems[i].Sku = helpers.GenerateItemSKU(product, p)

		err := s.FileStorageSrv.UploadImage(ctx, product.ProductItems[i].FilePath,
			product.ProductItems[i].File,
			product.ProductItems[i].ContentType)
		if err != nil {
			return err
		}

		product.ProductItems[i].Options = nil
	}

	err = s.ProductRepo.Insert(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
