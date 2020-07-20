package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetHeaderContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleRequests() {
	router := mux.NewRouter()

	router.Use(SetHeaderContentType)

	router.HandleFunc("/api/users", GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUserByID).Methods("GET")
	router.HandleFunc("/api/users/{name}/{email}", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")

	fmt.Println("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Migration()

	handleRequests()
}
