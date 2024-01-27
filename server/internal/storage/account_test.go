package storage

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func createTestAccount(t *testing.T) *mongo.InsertOneResult {
	args := CreateAccountParams{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 6),
	}

	account, err := testDbc.CreateAccount(args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestListAccount(t *testing.T) {
	a1 := createTestAccount(t)
	a2, err := testDbc.ListAccount(a1.InsertedID.(primitive.ObjectID).Hex())

	require.NoError(t, err)
	require.NotEmpty(t, a2)

	require.Equal(t, a1.InsertedID, a2.ID)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 5; i++ {
		createTestAccount(t)
	}

	_, err := testDbc.ListAccounts()
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	a1 := createTestAccount(t)
	err := testDbc.DeleteAccount(a1.InsertedID.(primitive.ObjectID).Hex())
	require.NoError(t, err)

	a2, err := testDbc.ListAccount(a1.InsertedID.(primitive.ObjectID).Hex())
	require.Error(t, err)
	require.EqualError(t, err, ErrAccountNotFound.Error())
	require.Empty(t, a2)
}
