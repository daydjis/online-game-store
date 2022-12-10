package main

import (
	"back/src/database"
	"errors"
	"regexp"
)

var ErrLoginLength = errors.New("login should contain from 8 to 24 symbols")
var ErrPasswordLength = errors.New("password should contain from 8 to 24 symbols")
var ErrForbiddenCharsLogin = errors.New("login contains forbidden characters")
var ErrForbiddenCharsPassword = errors.New("password contains forbidden characters")

func CheckLoginRequest(user database.User) error {
	// Проверяем длину логина
	if len(user.Login) < 8 || len(user.Login) > 24 {
		return ErrLoginLength
	}
	// Проверяем длину пароля
	if len(user.Password) < 8 || len(user.Password) > 24 {
		return ErrPasswordLength
	}
	// Проверяем символы логина
	matched, err := regexp.MatchString(`[^a-zA-z0-9_-]`, user.Login)
	if err != nil {
		return err
	} else if matched {
		return ErrForbiddenCharsLogin
	}
	// Проверяем символы пароля
	matched, err = regexp.MatchString(`[^a-zA-z0-9_@!-]`, user.Password)
	if err != nil {
		return err
	} else if matched {
		return ErrForbiddenCharsPassword
	}
	return err
}
