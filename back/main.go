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
		log.Println("GET /api/games")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Content-Type", "application/json")
		result := database.GetGames()
		err := json.NewEncoder(writer).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func newGamesHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		log.Println("POST /api/games/new")
		var v database.Game
		body, _ := io.ReadAll(request.Body)
		if err := json.Unmarshal(body, &v); err != nil {
			log.Fatal(err)
		}
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.Header().Set("Content-Type", "application/json")
		database.AddNewGame(v)
		err := json.NewEncoder(writer).Encode("{\"result\": \"success\"}")
		if err != nil {
			log.Fatal(err)
		}
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
