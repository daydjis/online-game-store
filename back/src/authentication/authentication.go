package authentication

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

const superSecretKey = "ILovePython"

func GenerateToken(login string) string {
	// Генерируем токен
	token := jwt.New(jwt.SigningMethodHS256)
	// Задаем параметры токена
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(12 * time.Hour)
	claims["authorized"] = true
	claims["login"] = login
	// Подписываем токен заданным ключом
	jwtToken, err := token.SignedString([]byte(superSecretKey))
	if err != nil {
		log.Fatal(err)
	}
	return jwtToken
}

func VerifyToken(clientToken string) {
	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}
