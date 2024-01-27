package storage

import (
	"errors"

	"github.com/oluijks/golang-starter/server/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dbc *Dbc) GetAccountByEmail(email string) (models.Account, error) {
	var account models.Account
	filter := bson.D{{Key: "email", Value: email}}
	err := dbc.GetAccountCollection().FindOne(dbc.ctx, filter).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Account{}, ErrAccountNotFound
		}
	}

	return account, nil
}
