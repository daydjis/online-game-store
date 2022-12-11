package authentication

import (
	"back/src/database"
	"errors"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

const superSecretKey = "ILovePython"

var ErrUnauthorized = errors.New("unauthorized")
var ErrUnableToExtractClaims = errors.New("unable to extract claims")

func GenerateToken(login string) string {
	// Генерируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt: time.Now().Unix()},
		Login: login,
	})
	// Подписываем токен заданным ключом
	jwtToken, err := token.SignedString([]byte(superSecretKey))
	if err != nil {
		log.Fatal(err)
	}
	return jwtToken
}

func VerifyToken(clientToken string) error {
	// Объявляем функцию внутри функции и делаем что-то непонятное =)
	// Парсим токен
	if len(clientToken) == 0 {
		return ErrUnauthorized
	}
	token, err := jwt.ParseWithClaims(clientToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(superSecretKey), nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// Вынимаем закодированный в токене логин и ищем в БД пользователя с таким логином
	claims, ok := token.Claims.(*tokenClaims)
	if ok && token.Valid {
		var user database.User
		user.Login = claims.Login
		userDB, errDB := database.GetUserByLogin(user)
		if errDB != nil || userDB.Login == "" {
			return ErrUnauthorized
		}
		return nil
	}
	return ErrUnableToExtractClaims
}
