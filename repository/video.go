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

func GetVideos(filter bson.M, duration time.Duration) []*model.Video {
	if env.MONGO_DATABASE == "" || env.MONGO_VIDEO_COLLECTION == "" {
		return make([]*model.Video, 0)
	}

	client := mongo.Client()
	if client == nil {
		return make([]*model.Video, 0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return make([]*model.Video, 0)
	}
	defer client.Disconnect(ctx)

	cursor, err := client.Database(env.MONGO_DATABASE).Collection(env.MONGO_VIDEO_COLLECTION).Find(ctx, filter)
	if err != nil {
		return make([]*model.Video, 0)
	}

	videos := make([]*model.Video, 0)
	if err = cursor.All(ctx, &videos); err != nil {
		return make([]*model.Video, 0)
	}

	for _, video := range videos {
		video.URL = storage.Sign(video.URL, duration)
	}

	return videos
}