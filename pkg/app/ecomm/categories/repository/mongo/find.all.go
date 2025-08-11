package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindAll(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("category mongo repo, GetAll failed: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var category domain.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, fmt.Errorf("category mongo repo, GetAll failed: %v", err)
		}
		categories = append(categories, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("category mongo repo, GetAll failed: %v", err)
	}

	return categories, nil
}
