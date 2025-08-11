package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) ConfirmEmail(ctx context.Context, userID string) error {
	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("userID invalid: %w", err)
	}

	filter := bson.M{
		"_id":            oID,
		"email_verified": false,
	}

	update := bson.M{
		"$set": bson.M{"email_verified": true},
	}

	result, err := r.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.ErrUserNotFound
	}

	return nil
}
