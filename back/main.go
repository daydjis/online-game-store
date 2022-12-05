package main

import (
	"back/src/database"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func gamesHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		result := database.GetGames()
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func newGamesHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var v database.Game
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &v); err != nil {
			log.Fatal(err)
		}
		database.AddNewGame(v)
	}
}

func main() {
	database.HealthCheck()
	http.HandleFunc("/api/games", gamesHandler)
	http.HandleFunc("/api/games/new", newGamesHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
