package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

// UserStruct
type User struct {
	gorm.Model
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;not null"`
}

var db *gorm.DB
var err error

// Migration GORM
func Migration() {
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")

	db, err := gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Println("Failed connecting to database")
	}
	db.AutoMigrate(&User{})

	defer db.Close()
}

// GetUsers getting all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User

	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")
	db, err := gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Println("Failed connecting to database")
	}
	defer db.Close()

	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

// CreateUser; create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")
	db, err := gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Println("Failed connecting to database")
	}
	defer db.Close()

	request := mux.Vars(r)
	name := request["name"]
	email := request["email"]

	newUser := db.Create(&User{Name: name, Email: email})

	json.NewEncoder(w).Encode(newUser)
}

// UpdateUser; Update one user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser; Delete one user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5432", "postgres", "gofirst", "26082002")
	db, err := gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Println("Failed connecting to database")
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var user User
	db.Table("users").Where("id = ?", id).Find(&user)
	deletedUser := db.Unscoped().Delete(&user)

	json.NewEncoder(w).Encode(deletedUser)
}
