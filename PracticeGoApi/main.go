package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range users {
		if strconv.Itoa(user.ID) == params["id"] {
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	users = append(users, User{ID: 1, Name: "Huda", Email: "huda@gmail.com"})
	users = append(users, User{ID: 2, Name: "Raihan", Email: "raihan@gmail.com"})

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUserById).Methods("GET")

	fmt.Println("Server started at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
