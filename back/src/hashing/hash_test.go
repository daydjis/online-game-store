package hashing

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mySecretPassword"
	hashedPassword := HashPassword(password)
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		t.Errorf(err.Error())
	}
}

func TestCheckPasswordPositive(t *testing.T) {
	password := "mySecretPassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err := CheckPassword(string(hashedPassword), password); err != nil {
		t.Errorf(err.Error())
	}
}

func TestCheckPasswordNegative(t *testing.T) {
	password := "mySecretPassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := CheckPassword(string(hashedPassword), "anotherPassword")
	if err.Error() != "wrong password" {
		t.Errorf(err.Error())
	}
}
