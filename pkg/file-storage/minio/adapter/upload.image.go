package adapter

import (
	"bytes"
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (a *Adapter) UploadImage(ctx context.Context, fileName string, file []byte, contentType string) error {
	
	fileReader := bytes.NewReader(file)

	options := minio.PutObjectOptions{
		ContentType: contentType,
	}

	fileSize := int64(len(file))
	_, err := a.MinioClient.PutObject(ctx, "ecomm", fileName, fileReader, fileSize, options)
	if err != nil {
		return fmt.Errorf("UploadImage err, %s", err.Error())
	}

	return nil
}
