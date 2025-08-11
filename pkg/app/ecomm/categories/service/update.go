package service

import "context"

func (s *Service) Update(ctx context.Context, id string, name string) error {
	return s.CategoryRepo.Update(ctx, id, name)
}
