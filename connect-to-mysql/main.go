package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("Connecting to MySQL Database...")

	// Make a connection to database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Connected to MySQL Database.")

	// Inserting data to table
	// insert, err := db.Query("INSERT INTO users VALUES('', 'Huda')")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("Successfully inserting a new data to users table")

	// Getting data from database
	results, err := db.Query("SELECT * FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.ID, &user.Name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.ID, user.Name)
	}

}
