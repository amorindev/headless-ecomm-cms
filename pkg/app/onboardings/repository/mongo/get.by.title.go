package mongo

import (
	"context"
	"errors"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
	obdErr "github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repository) GetByTitle(ctx context.Context, title string) (*domain.Onboarding, error) {
	var onboarding domain.Onboarding

	err := r.Collection.FindOne(ctx, bson.D{{Key: "title", Value: title}}).Decode(&onboarding)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, obdErr.ErrOnboardingNotFound
		}
		return nil, err
	}

	objID, ok := onboarding.ID.(bson.ObjectID)
	if !ok {
		return nil, errors.New("onboarding mongo repo - GetByTitle, failed to parse ID to ObjectID")
	}

	onboarding.ID = objID.Hex()
	return &onboarding, nil
}
