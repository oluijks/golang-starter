package storage

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/oluijks/golang-starter/server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testDbc *Dbc

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal(err)
	}
	clientOpts := options.Client().ApplyURI(config.MongoDBURL)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	testDbc = NewDbc(client, &config)
	os.Exit(m.Run())
}
