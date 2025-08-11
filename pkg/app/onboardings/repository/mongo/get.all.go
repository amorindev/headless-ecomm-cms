package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/onboardings/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) GetAll(ctx context.Context) ([]*domain.Onboarding, error) {
	var onboardings []*domain.Onboarding

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("onboarding mongo repo- Get: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var onboarding domain.Onboarding
		if err := cursor.Decode(&onboarding); err != nil {
			return nil, fmt.Errorf("onboarding mongo repo- Get: %v", err)
		}
		onboardings = append(onboardings, &onboarding)
	}
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("onboarding mongo repo- Get: %v", err)
	}
	return onboardings, nil
}
