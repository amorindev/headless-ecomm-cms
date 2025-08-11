package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.VariationRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewVariationRepo(client *mongo.Client, variationColl *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: variationColl,
	}
}
