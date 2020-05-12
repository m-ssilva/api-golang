package services

import (
	"io"
	"io/ioutil"
	"net/http"

	lib "github.com/m-ssilva/api-golang/lib"
)

// CreateUser get request body and call lib to parse and insert into database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	result, err := lib.CreateUser(body)
	if result != true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{ "code": 1, "message": "Internal Server Error" }`)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, `{ "success": true, "message": "User created" }`)
	}
}

// AuthenticateUser validates the request into database and authenticate it if is valid
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	result := lib.AuthenticateUser(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{ "success": true, "token":"`+result+`"}`)
}
