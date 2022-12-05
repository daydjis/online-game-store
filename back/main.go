package main

import (
	"back/src/database"
	"encoding/json"
	"log"
	"net/http"
)

func gamesHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		result := database.GetGames()
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	database.HealthCheck()
	http.HandleFunc("/api/games", gamesHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
