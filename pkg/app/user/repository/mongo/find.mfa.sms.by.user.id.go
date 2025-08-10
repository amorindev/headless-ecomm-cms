package mongo

import (
	"context"
	"errors"

	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
	userErr "github.com/amorindev/headless-ecomm-cms/pkg/app/user/errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindMfaSmsByUserID(ctx context.Context, userID string) (*domain.UserMfaSms, error) {
	userOID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	filter := bson.M{"user_id": userOID}

	var mfaSms domain.UserMfaSms
	err = r.TwoFaSmsColl.FindOne(ctx, filter).Decode(&mfaSms)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, userErr.ErrMfaSmsNotFound
		}
		return nil, fmt.Errorf("failed to find user MFA SMS by user ID: %w", err)
	}

	oID, ok := mfaSms.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("user mongo repo - FindMfaSmsByUserID, failed to convert ID to ObjectID")
	}
	mfaSms.ID = oID.Hex()

	phoneOID, ok := mfaSms.PhoneID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("user mongo repo - FindMfaSmsByUserID, failed to convert PhoneID to ObjectID")
	}
	mfaSms.PhoneID = phoneOID.Hex()

	mfaSms.UserID = userOID.Hex()

	return &mfaSms, nil
}
