package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/domain"
	userErr "github.com/amorindev/headless-ecomm-cms/pkg/app/users/errors"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) Find(ctx context.Context, id string) (*domain.User, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid type ID")
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	var user model.User
	err = r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, userErr.ErrUserNotFound
		}
		return nil, err
	}
	user.UserD.ID = id

	var roleIDs []string
	for _, oID := range user.RolesIDS {
		roleIDs = append(roleIDs, oID.Hex())
	}

	user.UserD.RolesIDs = roleIDs
	return user.UserD, nil
}
