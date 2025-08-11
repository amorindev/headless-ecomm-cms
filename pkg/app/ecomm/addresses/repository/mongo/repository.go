package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.AddressRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewAddressRepo(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}
