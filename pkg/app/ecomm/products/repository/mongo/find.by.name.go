package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	productErr "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindByName(ctx context.Context, name string) (*domain.Product, error) {
	var product domain.Product

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, productErr.ErrProductNotFound
		}

		return nil, err
	}

	oID, ok := product.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("product mongo repo - GetByName, failed to parse oID to ObjectID")
	}
	product.ID = oID.Hex()

	ctgOID, ok := product.CategoryID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("product mongo repo - GetByName, failed to parse ctgOID to ObjectID")
	}
	product.CategoryID = ctgOID.Hex()

	return &product, nil
}
