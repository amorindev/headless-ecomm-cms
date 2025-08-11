package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, phone *domain.Phone) error {
	id := bson.NewObjectID()
	phone.ID = id

	userOID, err := bson.ObjectIDFromHex(phone.UserID.(string))
	if err != nil {
		return errors.New("phone - failed to parse to objID")
	}

	phone.UserID = userOID

	_, err = r.Collection.InsertOne(context.Background(), phone)
	if err != nil {
		return fmt.Errorf("phone mongo repo: Insert err: %w", err)
	}
	phone.ID = id.Hex()
	phone.UserID = userOID.Hex()

	return nil
}
