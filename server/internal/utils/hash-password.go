package utils

import "golang.org/x/crypto/bcrypt"

func MakePasswordHash(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	return string(hashedPasswordBytes), err
}

func ComparePassword(hashedPassword, givenPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(givenPassword))
}
