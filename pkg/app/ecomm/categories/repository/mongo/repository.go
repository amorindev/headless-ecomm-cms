package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.CategoryRepo = &Repository{}

type Repository struct {
	Client *mongo.Client
	Collection   *mongo.Collection
}

func NewCategoryRepo(client *mongo.Client, categoryColl *mongo.Collection) *Repository{
	return &Repository{
		Client: client,
		Collection: categoryColl,
	}
}
