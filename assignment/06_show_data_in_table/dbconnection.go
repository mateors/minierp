package main

import (
	"database/sql"
	"log"
)

func dbcon() {
	// Connect to database
	db, err = sql.Open("sqlite3", "./minierp.db")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	log.Println("Database Connected....")
}
