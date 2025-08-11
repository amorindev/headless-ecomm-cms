package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
	varOptErr "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindByName(ctx context.Context, name string) (*domain.VariationOption, error) {
	var varOption domain.VariationOption

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&varOption)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, varOptErr.ErrVarOptNotFound
		}
		return nil, err
	}

	oID, ok := varOption.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("variation option mongo repo - GetByName, failed to parse ID to oID")
	}
	variationOID, ok := varOption.VariationID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("variation option mongo repo - GetByName, failed to parse variationOID to ObjectID")
	}

	varOption.ID = oID.Hex()
	varOption.VariationID = variationOID.Hex()

	return &varOption, nil
}
