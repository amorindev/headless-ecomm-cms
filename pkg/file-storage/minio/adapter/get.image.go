package adapter

import (
	"context"
	"net/url"
	"time"
)

func (a *Adapter) GetImage(ctx context.Context, fileName string) (string, error) {
	// * Set request parameters for content-disposition.
	reqParams := make(url.Values)

	// * Generates a presigned url which expires in a day.
	time := time.Hour * 24 * 7
	presignedURL, err := a.MinioClient.PresignedGetObject(ctx, "ecomm", fileName, time, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
