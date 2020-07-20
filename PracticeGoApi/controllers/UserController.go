package controllers

import (
	"encoding/json"
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

// GetAllUsers Getting all users with JSON
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users = append(users, User{ID: 1, Name: "Huda Prasetyo", Email: "huda@gmail.com"})
	users = append(users, User{ID: 2, Name: "Raihan", Email: "raihan@gmail.com"})

	json.NewEncoder(w).Encode(users)
}

// GetUserByID Get user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users = append(users, User{ID: 1, Name: "Huda Prasetyo", Email: "huda@gmail.com"})
	users = append(users, User{ID: 2, Name: "Raihan", Email: "raihan@gmail.com"})

	var params = mux.Vars(r)

	for _, user := range users {
		if strconv.Itoa(user.ID) == params["id"] {
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}
