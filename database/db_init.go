package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func DBInit() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/sea-salon.db")
	if err != nil {
		fmt.Println("Error opening database", err)
		return db
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS reviews (review_id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, rating TEXT, comment TEXT)")
	if err != nil {
		log.Fatalf("Error preparing create table statement: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("Error executing create table statement: %v", err)
	}
	statement, err = db.Prepare(`CREATE TABLE IF NOT EXISTS reservations (
		reservation_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		phone TEXT,
		type_of_service TEXT,
		date_and_time TEXT
	)`)
	if err != nil {
		log.Fatalf("Error preparing create table statement: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("Error executing create table statement: %v", err)
	}
	return db
}
