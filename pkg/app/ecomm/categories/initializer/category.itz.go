package initializer

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/port"
)

type CategoryItz struct {
	CategoryRepo port.CategoryRepo
}

func NewCategoryItz(categoryRepo port.CategoryRepo) *CategoryItz {
	return &CategoryItz{
		CategoryRepo: categoryRepo,
	}
}

func (ci *CategoryItz) SeedEssential(ctx context.Context) error {
	ctgs := []*domain.Category{
		domain.NewCategory("Men"),
		domain.NewCategory("Women"),
		domain.NewCategory("Unisex"),
		domain.NewCategory("boys and girls"),
	}

	for _, ctg := range ctgs {
		exists, err := ci.CategoryRepo.Exists(ctx, ctg.Name)
		if err != nil {
			return err
		}
		if !exists {
			if err := ci.CategoryRepo.Insert(ctx, ctg); err != nil {
				return err
			}
		}
	}

	return nil
}
