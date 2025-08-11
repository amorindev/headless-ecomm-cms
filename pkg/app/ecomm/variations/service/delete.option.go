package service

import "context"




func (s *Service) DeleteOption(ctx context.Context, varOptID string) error {
	return s.VarOptRepo.Delete(ctx, varOptID)
}
