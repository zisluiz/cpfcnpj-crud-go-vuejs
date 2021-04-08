package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection    *mongo.Collection
	clientOptions *options.ClientOptions
	databaseName  = os.Getenv("MONGODB_DATABASE_NAME")
)

func init() {
	clientOptions = options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION"))
}

func GetClient() (func() *mongo.Database, context.Context, func(), error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, nil, nil, err
	} else {
		database := func() *mongo.Database {
			var database = client.Database(databaseName)
			return database
		}

		close := func() {
			client.Disconnect(ctx)
		}

		return database, ctx, close, err
	}
}
