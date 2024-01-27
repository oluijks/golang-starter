package database

import (
	"context"
	"log"

	"github.com/oluijks/golang-starter/server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

func NewMongoConnection(lc fx.Lifecycle, config *config.Config) *mongo.Client {
	clientOpts := options.Client().ApplyURI(config.MongoDBURL)
	client, err := mongo.Connect(context.Background(), clientOpts)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// Check the connection
			err = client.Ping(context.Background(), readpref.Primary())
			if err != nil {
				return err
			}
			log.Println("Connected to MongoDB!")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			client.Disconnect(ctx)
			return nil
		},
	})

	return client
}
