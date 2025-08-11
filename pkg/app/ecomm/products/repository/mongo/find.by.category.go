package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindByCategory(ctx context.Context, categoryID string) ([]*domain.Product, error) {
	oID, err := bson.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"category_id": oID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find products by category: %w", err)
	}
	defer cursor.Close(ctx)

	var results []*domain.Product
	for cursor.Next(ctx) {
		var p domain.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, fmt.Errorf("failed to decode product: %w", err)
		}
		results = append(results, &p)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error while reading products: %w", err)
	}

	return results, nil
}
