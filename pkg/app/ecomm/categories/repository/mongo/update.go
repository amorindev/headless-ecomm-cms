package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Update(ctx context.Context, id string, name string) error {

	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("category mongo repo, Update failed: %s", err.Error())
	}

	update := bson.M{"$set": bson.M{"name": name}}

	_, err = r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return fmt.Errorf("category mongo repo, Update failed: %s", err.Error())
	}

	return err
}
