package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item[T ShapedItem | ShapelessItem] struct {
	ItemType T
}

type ShapedItem struct {
	Id      primitive.ObjectID `bson:"_id"`
	Type    string             `bson:"type"`
	Key     interface{}        `bson:"key"`
	Pattern []string           `bson:"pattern"`
	Result  interface{}        `bson:"result"`
}

type ShapelessItem struct {
	Id          primitive.ObjectID `bson:"_id"`
	Type        string             `bson:"type"`
	Ingredients []interface{}      `bson:"ingredients"`
	Result      interface{}        `bson:"result"`
}
