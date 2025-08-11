package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
	phoneErr "github.com/amorindev/headless-ecomm-cms/pkg/app/phones/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindDefault(ctx context.Context) (*domain.Phone, error) {
	var phone domain.Phone

	err := r.Collection.FindOne(ctx, bson.D{{Key: "is_default", Value: true}}).Decode(&phone)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, phoneErr.ErrPhoneNotFound
		}
		return nil, err
	}

	oID, ok := phone.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("phone mongo repo - GetDefault, failed to parse ID to ObjectID")
	}
	phone.ID = oID.Hex()

	userOID, ok := phone.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - GetDefault, failed to parse UserID to ObjectID")
	}
	phone.UserID = userOID.Hex()
	return &phone, nil
}
