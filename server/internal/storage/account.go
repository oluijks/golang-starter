package storage

import (
	"errors"
	"log"
	"time"

	"github.com/oluijks/golang-starter/server/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type CreateAccountParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dbc *Dbc) CreateAccount(params CreateAccountParams) (*mongo.InsertOneResult, error) {
	newAccount := &models.Account{
		Email:     params.Email,
		Password:  params.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := dbc.GetAccountCollection().InsertOne(dbc.ctx, newAccount)
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (dbc *Dbc) ListAccount(id string) (models.Account, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Account{}, err
	}

	var account models.Account
	filter := bson.D{{Key: "_id", Value: objectID}}
	err = dbc.GetAccountCollection().FindOne(dbc.ctx, filter).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Account{}, ErrAccountNotFound
		}
	}

	return account, nil
}

func (dbc *Dbc) ListAccounts() ([]*models.Account, error) {
	cursor, err := dbc.GetAccountCollection().Find(dbc.ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var accounts []*models.Account
	if err = cursor.All(dbc.ctx, &accounts); err != nil {
		log.Fatal(err)
	}
	return accounts, nil
}

func (dbc *Dbc) UpdateAccount(account *models.Account, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.D{
		bson.E{Key: "email", Value: account.Email},
		bson.E{Key: "updated_at", Value: time.Now()},
	}

	result, err := dbc.GetAccountCollection().UpdateByID(dbc.ctx, objectID, bson.M{"$set": update})

	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount == 0 {
		return ErrAccountNotFound
	}

	return nil
}

func (dbc *Dbc) DeleteAccount(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	deleteResult, err := dbc.GetAccountCollection().DeleteOne(dbc.ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return ErrAccountNotFound
	}
	return nil
}

func (dbc *Dbc) DeleteAccounts() error {
	return nil
}

func (dbc *Dbc) GetAccountCollection() *mongo.Collection {
	return dbc.db.Database(dbc.config.DatabaseName).Collection(dbc.config.AccountCollectionName)
}
