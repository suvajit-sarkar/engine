package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// User are players connected
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func connect() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "ELLIOT:PASS@tcp(192.168.0.141:3306)/gamerpg")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT id,username FROM user_info")
	for results.Next() {
		var tag User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Username)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Username)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}
