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

func GetVideos(filter bson.M, pagination structs.Pagination, duration time.Duration) []*model.Video {
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

	size := pagination.GetSize()
	skip := pagination.GetSkip()
	cursor, err := client.Database(env.MONGO_DATABASE).Collection(env.MONGO_VIDEO_COLLECTION).Find(ctx, filter, &options.FindOptions{
		Limit: &size,
		Skip:  &skip,
		Sort:  bson.M{"_id": 1},
	})
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