package mongo

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ port.OnboardingRepo = &Repository{}

type Repository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewOnboardingRepo(client *mongo.Client, collection *mongo.Collection) *Repository {
	return &Repository{
		Client:     client,
		Collection: collection,
	}
}
