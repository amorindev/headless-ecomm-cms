package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
	addressErr "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetDefault(ctx context.Context) (*domain.Address, error) {
	var address domain.Address

	err := r.Collection.FindOne(ctx, bson.D{{Key: "is_default", Value: true}}).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, addressErr.ErrAddressNotFound
		}
		return nil, err
	}

	oID, ok := address.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - GetDefault, failed to parse oID to ObjectID")
	}
	address.ID = oID.Hex()

	userOID, ok := address.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - GetDefault, failed to parse userOID to ObjectID")
	}
	address.UserID = userOID.Hex()
	return &address, nil
}
