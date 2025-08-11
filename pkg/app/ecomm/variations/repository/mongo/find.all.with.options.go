package mongo

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindAllWithOptions(ctx context.Context) ([]*domain.Variation, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "var_options"},
			{Key: "localField", Value: "_id"},
			{Key: "foreignField", Value: "variation_id"},
			{Key: "as", Value: "options"},
		}}},
	}

	cursor, err := r.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var variations []*domain.Variation
	if err := cursor.All(ctx, &variations); err != nil {
		return nil, err
	}

	return variations, nil
}
