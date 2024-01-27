package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := "S3cR@p@sSw0rD"

	hashedPasword, err := MakePasswordHash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPasword)

	err = ComparePassword(hashedPasword, password)
	require.NoError(t, err)

	incorrectPassword := "z4dS&pIZzv1rT"
	err = ComparePassword(hashedPasword, incorrectPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
