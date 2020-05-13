package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-ssilva/api-golang/middlewares"
	"github.com/m-ssilva/api-golang/services"
)

func main() {
	fmt.Println("Starting API...")
	router := mux.NewRouter()

	router.Handle("/register", middlewares.RootHandler(services.CreateUser)).Methods("POST")
	router.Handle("/login", middlewares.RootHandler(services.AuthenticateUser)).Methods("POST")
	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", router)
}
