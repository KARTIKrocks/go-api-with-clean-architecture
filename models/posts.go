package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty" bson:"title,omitempty"`
	Text  string             `json:"text,omitempty" bson:"text,omitempty"`
}
