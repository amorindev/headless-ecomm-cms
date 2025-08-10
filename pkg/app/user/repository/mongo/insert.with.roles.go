package mongo

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) InsertWithRoles(ctx context.Context, user *domain.User, rolesIDs []string) error {
	id := bson.NewObjectID()
	user.ID = id

	var rolesIDsInsert []bson.ObjectID
	for _, idStr := range rolesIDs {
		oID, err := bson.ObjectIDFromHex(idStr)
		if err != nil {
			return err
		}
		rolesIDsInsert = append(rolesIDsInsert, oID)
	}

	// * User Model from UserDomain
	model := model.User{
		UserD:    user,
		RolesIDS: rolesIDsInsert,
	}

	_, err := r.Collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}

	user.ID = id.Hex()

	return nil
}
