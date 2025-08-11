package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) GetAll(ctx context.Context, userID string) ([]*domain.Address, error) {
	oID, err := bson.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("address - invalid userID: %w", err)
	}
	filter := bson.M{"user_id": oID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("address - error finding addresses: %w", err)
	}
	defer cursor.Close(ctx)

	var addresses []*domain.Address
	if err := cursor.All(ctx, &addresses); err != nil {
		return nil, fmt.Errorf("address - error decoding addresses: %w", err)
	}

	return addresses, nil
}
