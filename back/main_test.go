package main

import (
	"back/src/database"
	"back/src/hashing"
	"errors"
	"fmt"
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

func TestMakeResponseForDelete(t *testing.T) {
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

type MakeResponseForRegisterParameters struct {
	login            string
	err              interface{ Error() string }
	expectedResponse string
	expectedStatus   int
}

var MakeResponseForRegisterValues = []MakeResponseForRegisterParameters{
	{"login", errors.New("such a horrible error"), "{\"Result\":\"User was not created,\"Error\":\"such a horrible error\"}", 500},
	{"login", database.ErrDuplicatedLogin, "{\"Result\":\"User was not created\",\"Error\":\"User with this login already exists\"}", 409},
}

func TestMakeResponseForRegister(t *testing.T) {
	for _, arg := range MakeResponseForRegisterValues {
		response, status := MakeResponseForRegister(arg.login, arg.err)
		if response != arg.expectedResponse {
			t.Errorf("Expected response: %s, Got: %s", arg.expectedResponse, response)
		}
		if status != arg.expectedStatus {
			t.Errorf("Expected status: %d, Got: %d", arg.expectedStatus, status)
		}

	}
}

type MakeResponseForLoginParameters struct {
	login            string
	err              error
	expectedResponse string
	expectedStatus   int
}

var MakeResponseForLoginValues = []MakeResponseForLoginParameters{
	{"login", hashing.ErrWrongPassword, fmt.Sprintf("{\"Result\":\"Unsuccessful login,\"Error\":\"%s\"}", hashing.ErrWrongPassword), 404},
	{"login", database.ErrWrongLogin, fmt.Sprintf("{\"Result\":\"Unsuccessful login,\"Error\":\"%s\"}", database.ErrWrongLogin), 404},
	{"login", errors.New("such a horrible error"), "{\"Result\":\"Unsuccessful login,\"Error\":\"such a horrible error\"}", 500},
}

func TestMakeResponseForLogin(t *testing.T) {
	for _, arg := range MakeResponseForLoginValues {
		response, status := MakeResponseForLogin(arg.login, arg.err)
		if response != arg.expectedResponse {
			t.Errorf("Expected response: %s, Got: %s", arg.expectedResponse, response)
		}
		if status != arg.expectedStatus {
			t.Errorf("Expected status: %d, Got: %d", arg.expectedStatus, status)
		}

	}
}

type CheckLoginParameters struct {
	database.User
	expectedError error
}

var CheckLoginValues = []CheckLoginParameters{
	{User: database.User{Login: "login"}, expectedError: ErrLoginLength},
	{User: database.User{Login: "loginWhichIsSuperExtremelyLongForValidation"}, expectedError: ErrLoginLength},
	{User: database.User{Login: "normal_login", Password: "pass"}, expectedError: ErrPasswordLength},
	{User: database.User{Login: "normal_login", Password: "passwordWhichIsSuperExtremelyLongForValidation"}, expectedError: ErrPasswordLength},
	{User: database.User{Login: "loginForbidden!", Password: "normal_pa55sword"}, expectedError: ErrForbiddenCharsLogin},
	{User: database.User{Login: "normal_login", Password: "passwordBad/"}, expectedError: ErrForbiddenCharsPassword},
}

func TestCheckLoginRequest(t *testing.T) {
	for _, arg := range CheckLoginValues {
		err := CheckLoginRequest(arg.User)
		if err != arg.expectedError {
			t.Errorf("Wrong error, got: %s, expected: %s", err, arg.expectedError.Error())
		}
	}
}
