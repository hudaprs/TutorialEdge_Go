package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hudaprs/TutorialEdge_Go/PracticeGoApi/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")

	fmt.Println("Server started at port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
