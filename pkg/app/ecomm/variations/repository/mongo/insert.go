package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, variation *domain.Variation) error {
	id := bson.NewObjectID()
	variation.ID = id

	_, err := r.Collection.InsertOne(context.Background(), variation)
	if err != nil {
		return fmt.Errorf("category mongo repo: Create err: %w", err)
	}
	variation.ID = id.Hex()
	return nil
}
