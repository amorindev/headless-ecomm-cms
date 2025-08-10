package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) UpdateMfaStatus(ctx context.Context, userID string, mfaStatus *domain.MfaStatus) error {
	userOID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{"mfa_status": mfaStatus}}
	resp, err := r.Collection.UpdateByID(ctx, userOID, update)
	if err != nil {
		return err
	}
	if resp.MatchedCount == 0 {
		return errors.New("mfa status not found")
	}
	return nil
}
