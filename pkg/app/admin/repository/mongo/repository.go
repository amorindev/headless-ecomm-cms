package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/admin/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.AdminRepo = &Repository{}

type Repository struct {
	Client   *mongo.Client
	UserColl *mongo.Collection
}

func NewAdminRepo(
	client *mongo.Client,
	userColl *mongo.Collection,
) *Repository {
	return &Repository{
		Client:   client,
		UserColl: userColl,
	}
}
