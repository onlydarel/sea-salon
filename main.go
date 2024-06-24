package main

import (
	"database/sql"
	"fmt"
	"github.com/onlydarel/sea-salon/database"
	"log"
	"net/http"
)

const port = ":8080"

var db *sql.DB

// main is the main function to run the application
func main() {
	db = database.DBInit()
	if db == nil {
		log.Println("Error initializing database")
	}
	fmt.Println("Successfully initialized database")
	fmt.Println("Listening on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(db),
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
