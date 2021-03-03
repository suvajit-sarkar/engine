package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/suvajit-sarkar/engine/world"
)

// WorldList contains slice of world list
type WorldList struct {
	List []*world.World `json:"list"`
}

func newDao(query string) ([]byte, error) {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/game")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//Create placeholder for the return object
	switch query {
	case "getWorldList":
		fmt.Println("Dao query for world list")
		daoResult := WorldList{[]*world.World{}}
		results, err := db.Query("SELECT * FROM world_data")
		for results.Next() {
			var tag world.World
			// for each row, scan the result into our tag composite object
			err = results.Scan(&tag.WorldID, &tag.WorldName, &tag.WorldType, &tag.WorldCreationTime, &tag.WorldPopulation)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			daoResult.List = append(daoResult.List, &tag)
		}
		return json.Marshal(daoResult)
	case "":
		fmt.Println("")
		daoResult := WorldList{[]*world.World{}}
		return json.Marshal(daoResult)
	default:
		fmt.Println("Sorry cant take default.")
		daoResult := WorldList{[]*world.World{}}
		return json.Marshal(daoResult)
	}

}
