package mongo

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindByVariationID(ctx context.Context, variationID string) ([]*domain.VariationOption, error) {
	objID, err := bson.ObjectIDFromHex(variationID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.Collection.Find(ctx, bson.M{"variation_id": objID})
	if err != nil {
		return nil, err
	}
	var opts []*domain.VariationOption

	err = cursor.All(ctx, &opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
