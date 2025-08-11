package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, otp *domain.OtpCodes) error {
	id := bson.NewObjectID()
	otp.ID = id

	userOID, err := bson.ObjectIDFromHex(otp.UserID.(string))
	if err != nil {
		return fmt.Errorf("auth mongo repo - Create err: %w", err)
	}
	otp.UserID = userOID

	_, err = r.Collection.InsertOne(context.Background(), otp)
	if err != nil {
		return fmt.Errorf("otp mongo repo: Create err: %w", err)
	}

	otp.ID = id.Hex()
	otp.UserID = userOID.Hex()

	return nil
}