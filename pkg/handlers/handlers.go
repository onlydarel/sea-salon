package handlers

import (
	"database/sql"
	"fmt"
	"github.com/onlydarel/sea-salon/models"
	"github.com/onlydarel/sea-salon/pkg/render"
	"log"
	"net/http"
)

var defaultReviews = []models.Review{
	{Name: "Felix Rafael", Rating: "5", Comment: "Salonnya bagus banget aku bisa jadi dirt block dengan pd 😘"},
	{Name: "Ady Wijaya", Rating: "1", Comment: "Jelekk banget aku jadi KEK TELOR mukanya -100/10 😡"},
	{Name: "Hans Christian", Rating: "4", Comment: "Bagus banget aku di rekomendasiin sama pelik makash beb 😍"},
}

// Home handles the home page
func Home(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Check if the form was submitted
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		// Create a new review from form data
		review := models.Review{
			Name:    r.Form.Get("name"),
			Rating:  getRatingValue(r),
			Comment: r.Form.Get("comment"),
		}
		// Insert the review into the database
		statement, err := db.Prepare("INSERT INTO reviews (name, rating, comment) VALUES (?, ?, ?)")
		if err != nil {
			log.Printf("Error preparing insert statement: %v\n", err)
			http.Error(w, "Error preparing insert statement", http.StatusInternalServerError)
			return
		}
		_, err = statement.Exec(review.Name, review.Rating, review.Comment)
		if err != nil {
			log.Printf("Error executing insert statement: %v\n", err)
			http.Error(w, "Error executing insert statement", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	rows, err := db.Query("SELECT review_id, name, rating, comment FROM reviews")
	if err != nil {
		log.Printf("Error querying database: %v\n", err)
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		err := rows.Scan(&review.ReviewID, &review.Name, &review.Rating, &review.Comment)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		reviews = append(reviews, review)
	}

	// If there are no reviews in the database, use the default reviews
	if len(reviews) == 0 {
		reviews = defaultReviews
	}

	data := map[string]interface{}{
		"Reviews": reviews,
	}
	render.RenderTemplate(w, "index.page.html", data)
}

// getRatingValue extracts the selected rating value from the form data
func getRatingValue(r *http.Request) string {
	// Loop through the radio button options to find the selected one
	ratings := []string{"1", "2", "3", "4", "5"}
	for _, rating := range ratings {
		if r.Form.Get("rating") == rating {
			return rating
		}
	}
	// Default to an empty string if no rating is selected
	return "3"
}

func Reservation(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		// Create a new review from form data
		new_reservation := models.Reservation{
			Name:          r.Form.Get("name"),
			TypeOfService: r.Form.Get("service_type"),
			DateAndTime:   r.Form.Get("date_and_time"),
			Phone:         r.Form.Get("phone_number"),
		}
		fmt.Printf("new_reservation: %v\n", new_reservation)
		// Insert the reservation into the database
		statement, err := db.Prepare("INSERT INTO reservations (name, phone, type_of_service, date_and_time) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Printf("Error preparing insert statement: %v\n", err)
			http.Error(w, "Error preparing insert statement", http.StatusInternalServerError)
			return
		}
		_, err = statement.Exec(new_reservation.Name, new_reservation.Phone, new_reservation.TypeOfService, new_reservation.DateAndTime)
		if err != nil {
			log.Printf("Error executing insert statement: %v\n", err)
			http.Error(w, "Error executing insert statement", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/reserve", http.StatusFound)
		return
	}
	render.RenderTemplate(w, "reservation.page.html", nil)
}

// Assets is a handler for static asset files in the templates folder
func Assets(w http.ResponseWriter, r *http.Request) {
	// Create the handler for serving files in the assets folder
	fileServer := http.FileServer(http.Dir("templates/assets"))
	http.StripPrefix("/assets/", fileServer).ServeHTTP(w, r)
}

// Styles is a handler for css files in the templates folder
func Styles(w http.ResponseWriter, r *http.Request) {
	// Create the handler for serving files in the styles folder
	fileServer := http.FileServer(http.Dir("templates/styles"))
	http.StripPrefix("/styles/", fileServer).ServeHTTP(w, r)
}
