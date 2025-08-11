package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.VariationOptionRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewVarOptRepo(client *mongo.Client, varOptColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: varOptColl,
	}
}
