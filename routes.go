package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/onlydarel/sea-salon/pkg/handlers"
	"net/http"
)

func routes(db *sql.DB) http.Handler {
	mux := chi.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(db, w, r)
	})
	mux.HandleFunc("/reserve", func(w http.ResponseWriter, r *http.Request) {
		handlers.Reservation(db, w, r)
	})
	mux.HandleFunc("/assets/*", handlers.Assets)
	mux.HandleFunc("/styles/*", handlers.Styles)

	return mux
}
