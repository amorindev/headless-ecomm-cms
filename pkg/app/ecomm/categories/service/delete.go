package service

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, id string) error {
	products, err := s.ProductRepo.FindByCategory(ctx, id)
	if err != nil {
		return err
	}

	if len(products) > 0 {
		return fmt.Errorf("the category cannot be eliminated, it has associated products")
	}

	return s.CategoryRepo.Delete(ctx, id)
}
