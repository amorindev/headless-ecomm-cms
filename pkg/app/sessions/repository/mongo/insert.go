package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/sessions/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, session *domain.Session) error {
	objID := bson.NewObjectID()
	session.ID = objID

	userOID, err := bson.ObjectIDFromHex(session.UserID.(string))
	if err != nil {
		return fmt.Errorf("session mongo repo - invalid user ID format: %w", err)
	}

	session.UserID = userOID

	_, err = r.Collection.InsertOne(ctx, session)
	if err != nil {
		return fmt.Errorf("session mongo repo - failed to insert session: %w", err)
	}

	session.ID = objID.Hex()
	session.UserID = userOID.Hex()

	return nil
}
