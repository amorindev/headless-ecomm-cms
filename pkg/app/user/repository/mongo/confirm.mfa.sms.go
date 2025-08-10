package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) ConfirmMfaSms(ctx context.Context, userID string) error {

	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("userID invalid: %w", err)
	}

	filter := bson.M{
		"user_id":   oID,
		"confirmed": false,
	}

	update := bson.M{
		"$set": bson.M{"confirmed": true},
	}

	result, err := r.TwoFaSmsColl.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no matching unconfirmed MFA-SMS record found for the given user")
	}

	return nil
}
