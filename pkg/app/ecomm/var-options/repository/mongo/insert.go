package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, varOption *domain.VariationOption) error {
	id := bson.NewObjectID()
	varOption.ID = id

	variationObjID, err := bson.ObjectIDFromHex(varOption.VariationID.(string))
	if err != nil {
	  return fmt.Errorf("varOption mongo repo - Create err %v", err)
	}
	varOption.VariationID = variationObjID

	_, err = r.Collection.InsertOne(ctx,varOption)
	if err != nil {
		return fmt.Errorf("varOption mongo repo - Create err %v", err)
	}
	
	varOption.ID = id.Hex()
	varOption.VariationID = variationObjID.Hex()
	return nil
}


