package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, onboarding *domain.Onboarding) error {
	id := bson.NewObjectID()
	onboarding.ID = id

	_, err := r.Collection.InsertOne(context.Background(), onboarding)
	if err != nil {
		return fmt.Errorf("onboarding mongo repo: Insert err: %w", err)
	}
	onboarding.ID = id.Hex()

	return nil
}
