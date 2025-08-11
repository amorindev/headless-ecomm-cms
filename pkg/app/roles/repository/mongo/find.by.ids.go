package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/roles/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) FindByIDs(ctx context.Context, roleIDs []string) ([]string, error) {

	var roleOIDs []bson.ObjectID
	for _, id := range roleIDs {
		oID, err := bson.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		roleOIDs = append(roleOIDs, oID)
	}

	if len(roleOIDs) == 0 {
		return nil, nil
	}

	filter := bson.M{"_id": bson.M{"$in": roleOIDs}}
    
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find roles: %w", err)
	}
	defer cursor.Close(ctx)

	var roleNames []string
	for cursor.Next(ctx) {
		var role domain.Role
		if err := cursor.Decode(&role); err != nil {
			return nil, fmt.Errorf("failed to decode role: %w", err)
		}
		roleNames = append(roleNames, role.Name)
	}

	return roleNames, nil
}
