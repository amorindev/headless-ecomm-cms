package mongo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
    "github.com/amorindev/headless-ecomm-cms/pkg/app/user/port"
)

var _ port.UserRepo = &Repository{}

type Repository struct {
	Client       *mongo.Client
	Collection   *mongo.Collection
	TwoFaSmsColl *mongo.Collection
}

func NewUserRepo(
	client *mongo.Client,
	collection *mongo.Collection,
	twoFaSmsColl *mongo.Collection,
) *Repository {
	return &Repository{
		Client:       client,
		Collection:   collection,
		TwoFaSmsColl: twoFaSmsColl,
	}
}
