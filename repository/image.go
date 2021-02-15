package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"turrium/env"
	"turrium/model"
	"turrium/mongo"
	"turrium/storage"
	"turrium/structs"
)

func GetImages(filter bson.M, pagination structs.Pagination, duration time.Duration) []*model.Image {
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

	size := pagination.GetSize()
	skip := pagination.GetSkip()
	cursor, err := client.Database(env.MONGO_DATABASE).Collection(env.MONGO_IMAGE_COLLECTION).Find(ctx, filter, &options.FindOptions{
		Limit: &size,
		Skip:  &skip,
		Sort:  bson.M{"_id": 1},
	})
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