package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/m-ssilva/api-golang/services"
)

func main() {
	fmt.Println("Starting API...")
	router := mux.NewRouter()

	router.HandleFunc("/register", services.CreateUser).Methods("POST")
	http.ListenAndServe(":8000", router)
}
