package service

import (
	"context"
)

func (s *Service) GetImage(ctx context.Context, fileName string) (string, error) {
	url, err := s.FileStgAdp.GetImage(ctx, fileName)
	if err != nil {
		return "", err
	}
	return url, err
}
