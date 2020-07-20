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
	router.HandleFunc("/api/users/{name}/{email", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Migration()

	handleRequests()
}
