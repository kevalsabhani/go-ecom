package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(password),
	)
	return err == nil
}
