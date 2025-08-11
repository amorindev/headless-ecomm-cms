package mongo

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindAll(ctx context.Context) ([]*domain.Variation, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var variations []*domain.Variation
	err = cursor.All(ctx, &variations)
	if err != nil {
		return nil, err
	}
	return variations, nil
}
