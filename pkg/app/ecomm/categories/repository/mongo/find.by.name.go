package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/domain"
	categoryErr "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindByName(ctx context.Context, name string) (*domain.Category, error) {
	var category domain.Category

	err := r.Collection.FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, categoryErr.ErrCategoryNotFound
		}
		return nil, err
	}

	oID, ok := category.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("category mongo repo - GetByName, failed to parse oID to ObjectID")
	}
	category.ID = oID.Hex()
	return &category, nil
}
