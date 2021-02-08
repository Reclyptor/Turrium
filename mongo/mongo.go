package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"turrium/env"
)

func Client() *mongo.Client {
	if env.MONGO_HOST == "" || env.MONGO_DATABASE == "" || env.MONGO_USERNAME == "" || env.MONGO_PASSWORD == "" {
		return nil
	}

	credentials := options.Credential{
		Username:                env.MONGO_USERNAME,
		Password:                env.MONGO_PASSWORD,
		PasswordSet:             true,
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(env.MONGO_HOST).SetAuth(credentials))
	if err != nil {
		return nil
	}

	return client
}

func ListDatabases() []string {
	client := Client()
	if client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return nil
	}
	defer client.Disconnect(ctx)

	databases , err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return nil
	}

	return databases
}

func ListCollections(database string) []string {
	client := Client()
	if client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err := client.Connect(ctx)
	if err != nil {
		return nil
	}
	defer client.Disconnect(ctx)

	collections, err := client.Database(database).ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return nil
	}

	return collections
}