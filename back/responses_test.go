package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestMakeResponseForPostPositive(t *testing.T) {
	gameId := "randomId"
	expResponse := fmt.Sprintf("{\"Result\":\"Game was created successfully\", \"id\": \"%s\"}", gameId)
	expStatus := 200

	response, status := MakeResponseForPost(gameId, nil)
	if response != expResponse {
		t.Errorf("Expected response: %s, Got: %s", expResponse, response)
	}
	if status != expStatus {
		t.Errorf("Expected status: %d, Got: %d", expStatus, status)
	}
}

func TestMakeResponseForPostNegative(t *testing.T) {
	gameId := "randomId2"
	err := errors.New("such a horrible error")
	expResponse := fmt.Sprintf("{\"Result\":\"Game was not created,\"Error\": \"%s\"}", err)
	expStatus := 500

	response, status := MakeResponseForPost(gameId, err)
	if response != expResponse {
		t.Errorf("Expected response: %s, Got: %s", expResponse, response)
	}
	if status != expStatus {
		t.Errorf("Expected status: %d, Got: %d", expStatus, status)
	}
}
