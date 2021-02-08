package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/env"
	"turrium/model"
	"turrium/mongo"
	"turrium/storage"
)

func GetImages(filter bson.M, duration time.Duration) []*model.Image {
	if env.MONGO_DATABASE == "" || env.MONGO_IMAGE_COLLECTION == "" {
		return make([]*model.Image, 0)
	}

	client := mongo.Client()
	if client == nil {
		return make([]*model.Image, 0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return make([]*model.Image, 0)
	}
	defer client.Disconnect(ctx)

	cursor, err := client.Database(env.MONGO_DATABASE).Collection(env.MONGO_IMAGE_COLLECTION).Find(ctx, filter)
	if err != nil {
		return make([]*model.Image, 0)
	}

	images := make([]*model.Image, 0)
	if err = cursor.All(ctx, &images); err != nil {
		return make([]*model.Image, 0)
	}

	for _, image := range images {
		image.URL = storage.Sign(image.URL, duration)
	}

	return images
}