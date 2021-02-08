package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Image struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Timestamp time.Time          `json:"timestamp"`
	Size      int                `json:"size"`
	Width     int                `json:"width"`
	Height    int                `json:"height"`
	Hidden    bool               `json:"hidden"`
	Type      string             `json:"type"`
	SHA1      string             `json:"sha1"`
	Filename  string             `json:"filename"`
	URL       string             `json:"url"`
	Caption   string             `json:"caption"`
	Tags      []string           `json:"tags"`
}