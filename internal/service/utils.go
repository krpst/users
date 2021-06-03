package service

import (
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd string) (hash string, err error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(h), err
}
