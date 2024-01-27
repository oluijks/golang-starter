package storage

import (
	"github.com/oluijks/golang-starter/server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Querier
}

type MongoDBStore struct {
	*Dbc
	client *mongo.Client
}

func NewDBStore(client *mongo.Client, config *config.Config) Store {
	return &MongoDBStore{
		client: client,
		Dbc:    NewDbc(client, config),
	}
}
