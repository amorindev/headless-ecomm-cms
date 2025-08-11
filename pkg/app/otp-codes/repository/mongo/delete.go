package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Delete(ctx context.Context, id string) error {

	otpObjID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("opt-codes mongo repo - Delete err: %w", err)
	}

	filter := bson.M{"_id": otpObjID}

	result, err := r.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("otp not found")
	}
	return nil
}
