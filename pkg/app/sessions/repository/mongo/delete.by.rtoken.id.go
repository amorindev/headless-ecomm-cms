package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) DeleteByRTokenID(ctx context.Context, rTokenID string) error {
    filter := bson.M{"refresh_token_id": rTokenID}

	_, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete refresh token by ID: %w", err)
	}

	return nil
}
