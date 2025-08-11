package mongo

import (
	"context"
	"errors"

	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
	addressErr "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Get(ctx context.Context, id string) (*domain.Address, error) {
	var address domain.Address

	oID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("address mongo repo - Get err: %w", err)
	}

	filter := bson.M{"_id": oID}

	err = r.Collection.FindOne(ctx, filter).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, addressErr.ErrAddressNotFound
		}
		return nil, fmt.Errorf("address mongo repo - Get err: %w", err)
	}

	address.ID = oID.Hex()

	userOID, ok := address.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - Get, failed to convert userOID to ObjectID")
	}
	address.UserID = userOID.Hex()

	storeOID, ok := address.StoreID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("address mongo repo - Get, failed to convert storeOID to ObjectID")
	}
	address.UserID = storeOID.Hex()

	return &address, nil
}
