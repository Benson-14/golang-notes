package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Println("Connected to SQLite successfully!")

	if err := db.Ping(); err != nil {
		log.Println("Ping failed:", err)
	} else {
		log.Println("Ping successful! DB is reachable.")
	}
}
