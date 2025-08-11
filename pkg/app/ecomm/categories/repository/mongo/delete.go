package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Delete(ctx context.Context, id string) error {

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("categories mongo repo, Delete failed: %s", err.Error())
	}

	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fmt.Errorf("categories mongo repo, Delete failed: %s", err.Error())
	}

	return err
}
