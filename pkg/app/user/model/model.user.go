package model

import (
	"github.com/amorindev/headless-ecomm-cms/pkg/app/user/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	UserD    *domain.User    `bson:"inline"`
	RolesIDS []bson.ObjectID `bson:"roles_ids"`
}
