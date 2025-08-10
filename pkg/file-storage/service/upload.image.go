package service

import (
	"context"
)

func (s *Service) UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error {
	err := s.FileStgAdp.UploadImage(ctx, fileName, file, contentType)
	if err != nil {
		return err
	}
	return nil
}
