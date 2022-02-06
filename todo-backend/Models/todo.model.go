package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDoItem struct {
	ID		primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Text	string				`json:"text" bson:"text"`
}