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
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Content-Type", "application/json")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusOK)
		return
	}
	if request.Method == http.MethodPost {
		var user database.User
		var response string
		var status int
		var token string
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
			// В случае отсутствия ошибок отправляем в ответе токен для аутентификации
			if err == nil {
				token = authentication.GenerateToken(user.Login)
			}
			response, status = MakeResponseForRegister(err, token)
		}
		// Формируем ответ клиенту
		writer.WriteHeader(status)
		if _, errResp := io.WriteString(writer, response); errResp != nil {
			log.Fatal(errResp)
		}
	}
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodOptions {
		writer.WriteHeader(http.StatusOK)
		return
	}
	if request.Method == http.MethodPost {
		var user database.User
		var response string
		var status int
		var token string
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
				response, status = MakeResponseForLogin(errDB, token)
			} else {
				// Если ошибка не возникла, сверяем пароль, указанный пользователем с паролем в БД
				errPassword := hashing.CheckPassword(userDB.Password, user.Password)
				if errPassword == nil {
					// В случае отсутствия ошибок отправляем в ответе токен для аутентификации
					token = authentication.GenerateToken(user.Login)
				}
				response, status = MakeResponseForLogin(errPassword, token)
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
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		token := request.Header.Get("Authorization")
		if token == "" {
			writer.WriteHeader(http.StatusForbidden)
			response := fmt.Sprint("{\"Error\":\"No Authorization header\"}")
			if _, errResp := io.WriteString(writer, response); errResp != nil {
				log.Fatal(errResp)
			}
			return
		}
		err := authentication.VerifyToken(token)
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
