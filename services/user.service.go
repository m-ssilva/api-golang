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

	_, err = lib.CreateUser(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{ "code": 1, "message": "Internal Server Error" }`)
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `{ "success": true, "message": "User created" }`)
}
