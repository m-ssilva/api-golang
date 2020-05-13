package services

import (
	"io"
	"io/ioutil"
	"net/http"

	lib "github.com/m-ssilva/api-golang/lib"
)

// CreateUser get request body and call lib to parse and insert into database
func CreateUser(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	err = lib.CreateUser(body)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{ "success": true, "message": "User created" }`)
	}

	return err
}

// AuthenticateUser validates the request into database and authenticate it if is valid
func AuthenticateUser(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	result, err := lib.AuthenticateUser(body)

	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{ "success": true, "token":"`+result+`"}`)
	}

	return err
}
