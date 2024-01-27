package storage

import (
	"github.com/oluijks/golang-starter/server/api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Querier interface {
	CreateAccount(params CreateAccountParams) (*mongo.InsertOneResult, error)
	DeleteAccount(id string) error
	DeleteAccounts() error
	GetAccountByEmail(email string) (models.Account, error)
	GetAccountCollection() *mongo.Collection
	ListAccount(id string) (models.Account, error)
	ListAccounts() ([]*models.Account, error)
	UpdateAccount(account *models.Account, id string) error
}

var _ Querier = (*Dbc)(nil)
