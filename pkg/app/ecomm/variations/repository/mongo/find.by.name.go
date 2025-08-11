package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
	variationErr "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindByName(ctx context.Context, name string) (*domain.Variation, error) {
	var variation domain.Variation

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&variation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, variationErr.ErrVariationNotFound
		}
		return nil, err
	}

	oID, ok := variation.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("variation mongo repo - GetByName, failed to parse ID to oID")
	}
	variation.ID = oID.Hex()
	return &variation, nil

}
