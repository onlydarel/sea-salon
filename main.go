package main

import (
	"fmt"
	"net/http"
	"progressNow/pkg/handlers"
)

const port = ":8080"

// main is the main function to run the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/assets/", handlers.Assets)
	http.HandleFunc("/styles/", handlers.Styles)

	fmt.Println("Listening on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}