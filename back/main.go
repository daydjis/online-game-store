package main

import (
	"back/src/database"
	"encoding/json"
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
	writer.Header().Set("Content-Type", "application/json")
	log.Println(request.Method, request.URL)
	if request.Method == http.MethodPost {
		var reqBody database.Game
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &reqBody); err != nil {
			log.Fatal(err)
		}
		result, err := database.AddNewGame(reqBody)
		response, status := MakeResponseForPost(result, err)
		writer.WriteHeader(status)
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
		deletedCount, err := database.DeleteGame(reqBody)
		response, status := MakeResponseForDelete(deletedCount, err)
		writer.WriteHeader(status)
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			log.Fatal(err)
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
			userID, err := database.RegisterNewUser(user)
			response, status = MakeResponseForRegister(userID, err)
		}
		// Формируем ответ клиенту
		writer.WriteHeader(status)
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
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
			// Если логин и пароль прошли валидацию, то авторизуем пользователя
			err := database.CheckLogin(user)
			response, status = MakeResponseForLogin(err)
		}
		// Формируем ответ клиенту
		writer.WriteHeader(status)
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	database.HealthCheck()
	http.HandleFunc("/api/games", getGamesHandler)
	http.HandleFunc("/api/games/new", createGameHandler)
	http.HandleFunc("/api/games/delete", deleteGameHandler)
	http.HandleFunc("/api/register", registrationHandler)
	http.HandleFunc("/api/login", loginHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
