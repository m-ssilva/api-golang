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

	result := lib.CreateUser(body)
	if result != true {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{ "code": 1, "message": "Internal Server Error" }`)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{ "success": true, "message": "User created" }`)
	}
}
