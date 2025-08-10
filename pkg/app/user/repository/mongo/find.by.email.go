package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
	userErr "github.com/amorindev/headless-ecomm-cms/pkg/app/user/errors"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user model.User

	err := r.Collection.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, userErr.ErrUserNotFound
		}
		return nil, err
	}

	objID, ok := user.UserD.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("user mongo repo - GetByEmail, failed to parse ID to ObjectID")
	}
	user.UserD.ID = objID.Hex()

	var roleIDs []string
	for _, oID := range user.RolesIDS {
		roleIDs = append(roleIDs, oID.Hex())
	}

	user.UserD.RolesIDs = roleIDs

	return user.UserD, nil
}
