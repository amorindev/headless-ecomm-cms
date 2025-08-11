package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, category *domain.Category) error {
	id := bson.NewObjectID()
	category.ID = id

	_, err := r.Collection.InsertOne(context.Background(), category)
	if err != nil {
		return fmt.Errorf("category mongo repo, Insert failed: %w", err)
	}
	category.ID = id.Hex()

	return nil
}
