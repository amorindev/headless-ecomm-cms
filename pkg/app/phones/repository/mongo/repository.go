package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.PhoneRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewPhoneRepository(client *mongo.Client, coll *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: coll,
	}
}
