package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Video struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Timestamp time.Time          `json:"timestamp"`
	Size      int                `json:"size"`
	Length    int                `json:"length"`
	Hidden    bool               `json:"hidden"`
	Type      string             `json:"type"`
	SHA1      string             `json:"sha1"`
	Filename  string             `json:"filename"`
	URL       string             `json:"url"`
	Title     string             `json:"title"`
	Tags      []string           `json:"tags"`
}