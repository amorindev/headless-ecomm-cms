package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/domain"
	otpErr "github.com/amorindev/headless-ecomm-cms/pkg/app/otp-codes/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Find(ctx context.Context, otpID string) (*domain.OtpCodes, error) {
	var otp domain.OtpCodes

	oID, err := bson.ObjectIDFromHex(otpID)
	if err != nil {
		return nil, errors.New("otp-codes mongo repo - invalid OTP ID format")
	}

	err = r.Collection.FindOne(ctx, bson.D{{Key: "_id", Value: oID}}).Decode(&otp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, otpErr.ErrOtpNotFound
		}
		return nil, fmt.Errorf("otp-codes mongo repo - failed to find OTP: %w", err)
	}

	userOID, ok := otp.UserID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("otp-codes mongo repo - failed to parse user ID as ObjectID")
	}

	otp.UserID = userOID.Hex()
	otp.ID = otpID
	return &otp, nil
}
