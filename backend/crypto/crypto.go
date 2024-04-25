package crypto

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RandSessionId() string {
	return uuid.New().String()
}

func MustHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
