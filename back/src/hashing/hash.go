package hashing

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrWrongPassword = errors.New("wrong password")

func HashPassword(password string) string {
	// Хэшируем пароль пользователя
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func CheckPassword(hashedPassword string, password string) error {
	// Сравниваем пароль, переданный пользователем, с хэшем пароля для данного пользователя
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errors.Is(bcrypt.ErrMismatchedHashAndPassword, err) {
		return ErrWrongPassword
	}
	return err
}
