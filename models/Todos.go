package Todos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Note    string             `json:"note,omitempty"`
	Checked bool               `json:"checked,omitempty"`
}
