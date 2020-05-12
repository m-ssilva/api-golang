package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-ssilva/api-golang/services"
)

func main() {
	fmt.Println("Starting API...")
	router := mux.NewRouter()

	router.HandleFunc("/register", services.CreateUser).Methods("POST")
	router.HandleFunc("/login", services.AuthenticateUser).Methods("POST")
	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", router)
}
