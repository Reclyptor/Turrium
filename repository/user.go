package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/env"
	"turrium/model"
	"turrium/mongo"
)

func GetUsers(filter bson.M) []*model.User {
	client := mongo.Client()
	if client == nil {
		return make([]*model.User, 0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return make([]*model.User, 0)
	}
	defer client.Disconnect(ctx)

	cursor, err := client.Database(env.MONGO_DATABASE).Collection(env.MONGO_USER_COLLECTION).Find(ctx, filter)
	if err != nil {
		return make([]*model.User, 0)
	}

	users:= make([]*model.User, 0)
	if err = cursor.All(ctx, &users); err != nil {
		return make([]*model.User, 0)
	}

	return users
}

func GetUserEmails(filter bson.M) map[string]struct{} {
	users := GetUsers(filter)
	emails := map[string]struct{}{}
	for _, user := range users {
		emails[user.Email] = struct{}{}
	}
	return emails
}