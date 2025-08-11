package mongo

import (
	"context"
	"fmt"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *Repository) Insert(ctx context.Context, product *domain.Product) error {
	// * Create ID
	id := bson.NewObjectID()
	product.ID = id

	for _, pItem := range product.ProductItems {
		pItemID := bson.NewObjectID()
		pItem.ID = pItemID
	}

	for _, pItem := range product.ProductItems {
		var ids []interface{}
		for _, varOtpID := range pItem.VarOptionIDs {
			oID, err := bson.ObjectIDFromHex(varOtpID.(string))
			if err != nil {
				return err
			}
			ids = append(ids, oID)
		}
		pItem.VarOptionIDs = ids
	}

	// * Assign ID category
	ctgOID, err := bson.ObjectIDFromHex(product.CategoryID.(string))
	if err != nil {
		return fmt.Errorf("product mongo repo - insert ctgOID from hex err %w", err)
	}

	product.CategoryID = ctgOID

	// * Insert in the database
	_, err = r.Collection.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("product mongo repo - Insert InsertOne err %w", err)
	}

	product.ID = id.Hex()
	product.CategoryID = ctgOID.Hex()

	for _, product := range product.ProductItems {
		product.ProductID = id.Hex()
	}

	return nil
}
