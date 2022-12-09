package main

import (
	"fmt"
	"log"
)

func MakeResponseForPost(gameId string, err error) (string, int) {
	var response string
	var status int
	if err != nil {
		log.Println(err)
		response = fmt.Sprintf("{\"Result\":\"Game was not created,\"Error\": \"%s\"}", err)
		status = 500
	} else {
		response = fmt.Sprintf("{\"Result\":\"Game was created successfully\", \"id\": \"%s\"}", gameId)
		status = 200
	}
	return response, status
}

func MakeResponseForDelete(deletedCount int64, err error) (string, int) {
	var response string
	var status int
	if err != nil {
		log.Println(err)
		response = fmt.Sprintf("{\"Result\":\"Game was not deleted,\"Error\": \"%s\"}", err)
		status = 500
	} else if deletedCount == 0 {
		response = fmt.Sprint("{\"Result\":\"Game was not deleted,\"Error\": \"Wrong game id\"}")
		status = 404
	} else {
		response = fmt.Sprint("{\"Result\":\"Game was deleted successfully\"}")
		status = 200
	}
	return response, status
}

func MakeResponseForRegister(userID string, err error) (string, int) {
	if err != nil {
		log.Println(err)
		if err.Error() == "duplicated login" {
			response := fmt.Sprint("{\"Result\":\"User was not created,\"Error\":\"User with this login already exists\"}")
			return response, 409
		}
		response := fmt.Sprintf("{\"Result\":\"User was not created,\"Error\":\"%s\"}", err)
		return response, 500
	} else {
		response := fmt.Sprintf("{\"Result\":\"User was created successfully\",\"id\": \"%s\"}", userID)
		return response, 200
	}
}
