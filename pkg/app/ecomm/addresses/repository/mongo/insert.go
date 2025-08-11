package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, address *domain.Address) error {
	id := bson.NewObjectID()
	address.ID = id

	userOID, err := bson.ObjectIDFromHex(address.UserID.(string))
	if err != nil {
		return errors.New("address- failed to parse to userOID")
	}

	address.UserID = userOID

	storeOID, err := bson.ObjectIDFromHex(address.StoreID.(string))
	if err != nil {
		return errors.New("address- failed to parse to storeOID")
	}

	address.StoreID = storeOID

	_, err = r.Collection.InsertOne(context.Background(), address)
	if err != nil {
		return fmt.Errorf("address mongo repo: Insert err: %w", err)
	}

	address.ID = id.Hex()
	address.UserID = userOID.Hex()
	address.StoreID = storeOID.Hex()

	return nil
}
