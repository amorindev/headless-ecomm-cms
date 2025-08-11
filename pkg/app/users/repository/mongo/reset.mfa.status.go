package mongo

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) ResetMfaStatus(ctx context.Context, userID string, mfaStatus *domain.MfaStatus) error {
	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oID}

	update := bson.M{
		"$set": bson.M{
			"mfa_status": mfaStatus,
		},
	}
	_, err = r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
