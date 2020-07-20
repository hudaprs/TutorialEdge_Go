package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{name}/{email}", CreateUser).Methods("POST")

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Migration()

	handleRequests()
}
