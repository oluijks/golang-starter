package storage

import (
	"context"

	"github.com/oluijks/golang-starter/server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dbc struct {
	db     *mongo.Client
	ctx    context.Context
	config *config.Config
}

func NewDbc(db *mongo.Client, config *config.Config) *Dbc {
	return &Dbc{db: db, config: config}
}
