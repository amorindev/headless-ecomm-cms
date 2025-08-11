package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Exists(ctx context.Context, name string) (bool, error) {
	filter := bson.M{"name": name}
	count, err := r.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
