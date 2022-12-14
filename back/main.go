package main

import (
	"back/src/authentication"
	"back/src/database"
	"back/src/hashing"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getGamesHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "application/json")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodGet {
		gameId := request.URL.Query().Get("id")
		result := database.GetGames(gameId)
		err := json.NewEncoder(writer).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createGameHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Content-Type", "application/json")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodPost {
		var game database.Game
		var response string
		var status int
		// Записываем тело запроса в переменную game
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &game); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("{\"Error\":\"%s\"}", err)
			if _, errResp := io.WriteString(writer, response); errResp != nil {
				log.Fatal(errResp)
			}
			return
		}
		// Проверяем корректность переданных параметров игры
		err := CheckCreateGameRequest(game)
		if err != nil {
			// Если параметры игры некорректны, возвращаем клиенту ошибку
			response, status = MakeResponseForClientError(err)
		} else {
			// Если параметры игры прошли валидацию, то создаем запись в БД
			result, err := database.AddNewGame(game)
			response, status = MakeResponseForPost(result, err)
		}
		// Формируем ответ клиенту
		writer.WriteHeader(status)
		if _, errResp := io.WriteString(writer, response); errResp != nil {
			log.Fatal(errResp)
		}
	}
}

func deleteGameHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "application/json")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodDelete {
		var reqBody database.Game
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &reqBody); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("{\"Error\":\"%s\"}", err)
			if _, errResp := io.WriteString(writer, response); errResp != nil {
				log.Fatal(errResp)
			}
			return
		}
		deletedCount, err := database.DeleteGame(reqBody)
		response, status := MakeResponseForDelete(deletedCount, err)
		writer.WriteHeader(status)
		if _, errResp := io.WriteString(writer, response); errResp != nil {
			log.Fatal(errResp)
		}
	}
}

func registrationHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "application/json")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodPost {
		var user database.User
		var response string
		var status int
		// Записываем тело запроса в переменную user
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &user); err != nil {
			log.Fatal(err)
		}
		// Проверяем корректность логина и пароля
		err := CheckLoginRequest(user)
		if err != nil {
			// Если логин и пароль не прошли валидацию, возвращаем ошибку и 404
			response, status = MakeResponseForClientError(err)
		} else {
			// Если логин и пароль прошли валидацию, то регистрируем пользователя
			_, err := database.RegisterNewUser(user)
			// В случае отсутствия ошибок отправляем в ответе куки для аутентификации
			if err == nil {
				token := authentication.GenerateToken(user.Login)
				http.SetCookie(writer, &http.Cookie{
					Name:    "auth_cookie",
					Value:   token,
					Expires: time.Now().Add(12 * time.Hour),
				})
			}
			response, status = MakeResponseForRegister(err)
		}
		// Формируем ответ клиенту
		writer.WriteHeader(status)
		if _, errResp := io.WriteString(writer, response); errResp != nil {
			log.Fatal(errResp)
		}
	}
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodPost {
		var user database.User
		var response string
		var status int
		// Записываем тело запроса в переменную user
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &user); err != nil {
			log.Fatal(err)
		}
		// Проверяем корректность логина и пароля
		err := CheckLoginRequest(user)
		if err != nil {
			// Если логин и пароль не прошли валидацию, возвращаем ошибку и 404
			response, status = MakeResponseForClientError(err)
		} else {
			// Если логин и пароль прошли валидацию, то получаем пользователя с указанным логином из базы
			userDB, errDB := database.GetUserByLogin(user)
			if errDB != nil {
				// Если возникла ошибка, например такого логина в базе нет, возвращаем ошибку
				response, status = MakeResponseForLogin(errDB)
			} else {
				// Если ошибка не возникла, сверяем пароль, указанный пользователем с паролем в БД
				errPassword := hashing.CheckPassword(userDB.Password, user.Password)
				if errPassword == nil {
					// В случае отсутствия ошибок отправляем в ответе куки для аутентификации
					token := authentication.GenerateToken(user.Login)
					http.SetCookie(writer, &http.Cookie{
						Name:    "auth_cookie",
						Value:   token,
						Expires: time.Now().Add(12 * time.Hour),
					})
				}
				response, status = MakeResponseForLogin(errPassword)
			}
		}
		// Формируем ответ клиенту
		writer.WriteHeader(status)
		if _, errResp := io.WriteString(writer, response); errResp != nil {
			log.Fatal(errResp)
		}
	}
}

func CheckToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		cookie, err := request.Cookie("auth_cookie")
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusForbidden)
			response := fmt.Sprintf("{\"Error\":\"%s\"}", err)
			if _, errResp := io.WriteString(writer, response); errResp != nil {
				log.Fatal(errResp)
			}
			return
		}
		err = authentication.VerifyToken(cookie.Value)
		if err != nil {
			writer.WriteHeader(http.StatusForbidden)
			response := fmt.Sprintf("{\"Error\":\"%s\"}", err)
			if _, errResp := io.WriteString(writer, response); errResp != nil {
				log.Fatal(errResp)
			}
			return
		}
		handler.ServeHTTP(writer, request)
	})
}

func main() {
	database.HealthCheck()
	http.HandleFunc("/api/games", getGamesHandler)
	http.Handle("/api/games/new", CheckToken(http.HandlerFunc(createGameHandler)))
	http.Handle("/api/games/delete", CheckToken(http.HandlerFunc(deleteGameHandler)))
	http.HandleFunc("/api/register", registrationHandler)
	http.HandleFunc("/api/login", loginHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
