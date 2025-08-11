package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) InsertMfaSms(ctx context.Context, twoFaSms *domain.UserMfaSms) error {
	id := bson.NewObjectID()
	twoFaSms.ID = id

	userObjID, err := bson.ObjectIDFromHex(twoFaSms.UserID.(string))
	if err != nil {
		return fmt.Errorf("mfa sms repo: invalid user_id format: %w", err)
	}

	phoneObjID, err := bson.ObjectIDFromHex(twoFaSms.PhoneID.(string))
	if err != nil {
		return fmt.Errorf("mfa sms repo: invalid phone_id format: %w", err)
	}
	twoFaSms.UserID = userObjID
	twoFaSms.PhoneID = phoneObjID

	_, err = r.TwoFaSmsColl.InsertOne(context.Background(), twoFaSms)
	if err != nil {
        return fmt.Errorf("mfa sms repo: failed to insert document: %w", err)
	}
	twoFaSms.ID = id.Hex()
	twoFaSms.UserID = userObjID.Hex()
	twoFaSms.PhoneID = phoneObjID.Hex()

	return nil
}
