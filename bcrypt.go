package leialib

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}

func BcryptValidate(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
