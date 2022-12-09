package main

import (
	"errors"
	"testing"
)

type MakeResponseForPostParameters struct {
	gameID           string
	err              interface{ Error() string }
	expectedResponse string
	expectedStatus   int
}

var MakeResponseForPostValues = []MakeResponseForPostParameters{
	{"randomId", nil, "{\"Result\":\"Game was created successfully\", \"id\": \"randomId\"}", 200},
	{"randomId2", errors.New("such a horrible error"), "{\"Result\":\"Game was not created,\"Error\": \"such a horrible error\"}", 500},
}

func TestMakeResponseForPost(t *testing.T) {
	for _, arg := range MakeResponseForPostValues {
		response, status := MakeResponseForPost(arg.gameID, arg.err)
		if response != arg.expectedResponse {
			t.Errorf("Expected response: %s, Got: %s", arg.expectedResponse, response)
		}
		if status != arg.expectedStatus {
			t.Errorf("Expected status: %d, Got: %d", arg.expectedStatus, status)
		}
	}

}

type MakeResponseForDeleteParameters struct {
	deletedCount     int64
	err              interface{ Error() string }
	expectedResponse string
	expectedStatus   int
}

var MakeResponseForDeleteValues = []MakeResponseForDeleteParameters{
	{0, errors.New("such a horrible error"), "{\"Result\":\"Game was not deleted,\"Error\": \"such a horrible error\"}", 500},
	{0, nil, "{\"Result\":\"Game was not deleted,\"Error\": \"Wrong game id\"}", 404},
	{1, nil, "{\"Result\":\"Game was deleted successfully\"}", 200},
}

func TestMakeResponseForDeleteErr(t *testing.T) {
	for _, arg := range MakeResponseForDeleteValues {
		response, status := MakeResponseForDelete(arg.deletedCount, arg.err)
		if response != arg.expectedResponse {
			t.Errorf("Expected response: %s, Got: %s", arg.expectedResponse, response)
		}
		if status != arg.expectedStatus {
			t.Errorf("Expected status: %d, Got: %d", arg.expectedStatus, status)
		}

	}
}
