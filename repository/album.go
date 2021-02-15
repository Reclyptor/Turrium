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

func GetAlbums(filter bson.M, pagination structs.Pagination, duration time.Duration) []*model.Album {
	client := mongo.Client()
	if client == nil {
		return make([]*model.Album, 0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return make([]*model.Album, 0)
	}
	defer client.Disconnect(ctx)

	size := pagination.GetSize()
	skip := pagination.GetSkip()
	cursor, err := client.Database(env.MONGO_DATABASE).Collection(env.MONGO_ALBUM_COLLECTION).Find(ctx, filter, &options.FindOptions{
		Limit: &size,
		Skip:  &skip,
		Sort:  bson.M{"_id": 1},
	})
	if err != nil {
		return make([]*model.Album, 0)
	}

	albums := make([]*model.Album, 0)
	if err = cursor.All(ctx, &albums); err != nil {
		return make([]*model.Album, 0)
	}

	for _, album := range albums {
		for _, image := range album.Images {
			image.URL = storage.Sign(image.URL, duration)
		}
		for _, video := range album.Videos {
			video.URL = storage.Sign(video.URL, duration)
		}
	}

	return albums
}