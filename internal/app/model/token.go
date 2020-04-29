package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Token struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Key         string                 `bson:"key" json:"key"`
	Description string                 `bson:"description" json:"description"`
	Active      bool                   `bson:"active" json:"active"`
	Params      map[string]interface{} `bson:"params" json:"params"`
}
