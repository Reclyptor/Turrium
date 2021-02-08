package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name        string           `json:"name"`
	Hidden      bool             `json:"hidden"`
	Description string           `json:"description"`
	Tags        []string         `json:"tags"`
	Images      []Image          `json:"images"`
	Videos      []Video          `json:"videos"`
}