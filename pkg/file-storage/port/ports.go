package port

import (
	"context"
)

type FileStorageAdapter interface {
	UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error
	GetImage(ctx context.Context, fileName string) (string, error)
}

type FileStorageSrv interface {
	UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error
	GetImage(ctx context.Context, fileName string) (string, error)
}
