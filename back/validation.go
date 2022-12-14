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

var ErrTitleLength = errors.New("title should contain from 3 to 24 symbols")
var ErrDescriptionLength = errors.New("description should contain at least 10 symbols")
var ErrPrice = errors.New("price should be greater than 0 and less than 50000")
var ErrGenresQuantity = errors.New("genres quantity should be greater than 1 and less than 8")
var ErrGenreName = errors.New("genre should contain from 4 to 20 symbols")
var ErrImage = errors.New("looks like your base64 encoded image is broken")
var ErrVideo = errors.New("looks like your video link is broken")
var ErrImageDescription = errors.New("image description should contain from 5 to 20 symbols")

func CheckCreateGameRequest(game database.Game) error {
	// Проверяем длину названия игры
	if len([]rune(game.Title)) < 3 || len([]rune(game.Title)) > 24 {
		return ErrTitleLength
	}
	// Проверяем длину поля описание игры
	if len([]rune(game.Description)) < 10 {
		return ErrDescriptionLength
	}
	// Проверяем цену игры
	if game.Price < 0 || game.Price > 50000 {
		return ErrPrice
	}
	// Проверяем количество жанров
	if len(game.Genres) < 2 || len(game.Genres) > 8 {
		return ErrGenresQuantity
	}
	// Проверяем указанные жанры
	for _, genre := range game.Genres {
		if len([]rune(genre)) < 3 || len([]rune(genre)) > 20 {
			return ErrGenreName
		}
	}
	// Проверяем строку с картинкой
	if len([]rune(game.Image)) < 50 {
		return ErrImage
	}
	// Проверяем строку с видео
	if len([]rune(game.Video)) < 5 {
		return ErrVideo
	}
	// Проверяем описание картинки
	if len([]rune(game.ImageDescription)) < 5 || len([]rune(game.ImageDescription)) > 20 {
		return ErrImageDescription
	}
	return nil
}
