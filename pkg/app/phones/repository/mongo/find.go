package mongo

import (
	"context"
	"errors"

	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
	errPhone "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Find(ctx context.Context, id string) (*domain.Phone, error) {
	var phone domain.Phone

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("phone mongo repo - Get err: %w", err)
	}

	filter := bson.M{"_id": objID}

	err = r.Collection.FindOne(ctx, filter).Decode(&phone)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errPhone.ErrPhoneNotFound
		}
		return nil, fmt.Errorf("phone mongo repo - Get err: %w", err)
	}

	phone.ID = objID.Hex()

	userObjID, ok := phone.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("phone mongo repo - Get, failed to convert UserID to ObjectID")
	}
	phone.UserID = userObjID.Hex()

	return &phone, nil
}
