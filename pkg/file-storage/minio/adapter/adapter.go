package adapter

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/file-storage/port"
	"github.com/minio/minio-go/v7"
)

var _ port.FileStorageAdapter = &Adapter{}

type Adapter struct {
	MinioClient *minio.Client
}

func NewAdapter(client *minio.Client) *Adapter {
	return &Adapter{
		MinioClient: client,
	}
}
