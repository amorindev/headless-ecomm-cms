package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Find(ctx context.Context, id string) (*domain.Product, error) {

	oID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID format: %w", err)
	}

	filter := bson.D{{Key: "_id", Value: oID}}

	var product domain.Product

	err = r.Collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrProductItemNotFound
		}
		return nil, err
	}

	product.ID = oID.Hex()
	return &product, nil
}