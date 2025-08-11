package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindAllByUserID(ctx context.Context, userID string) ([]*domain.Phone, error) {

	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("phone - invalid userID: %w", err)
	}
	filter := bson.M{"user_id": oID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("phone - error finding addresses: %w", err)
	}
	defer cursor.Close(ctx)

	var phones []*domain.Phone

	if err := cursor.All(ctx, &phones); err != nil {
		return nil, fmt.Errorf("phone - error decoding addresses: %w", err)
	}

	return phones, nil
}
