package minio

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
}

func NewClient() *MinioClient {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretKeyID := os.Getenv("MINIO_SECRET_KEY")
	/* useSSL := os.Getenv("MINIO_SECURE")

	var useSSLbool bool
	if useSSL == "true" {
		useSSLbool = true
	} else {
		useSSLbool = false
	} */

	// * Initialize minio client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretKeyID, ""),
		Secure: true,
	})

	if err != nil {
		log.Fatal("Minio client failed: %w", err)
	}

	newMinioClient := &MinioClient{
		Client: minioClient,
	}

	return newMinioClient
}

func (client *MinioClient) CreateStorage() {
	bucketName := os.Getenv("MINIO_BUCKET_NAME")
	if bucketName == "" {
		log.Fatal("environment variable MINIO_BUCKET_NAME is not set")
	}

	found, _ := client.Client.BucketExists(context.Background(), bucketName)
	if !found {
		err := client.Client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("create Storage failed: %s\n", err.Error())
		}
	}
}
