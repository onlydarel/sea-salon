package handlers

import (
	"fmt"
	"github.com/onlydarel/sea-salon/models"
	"github.com/onlydarel/sea-salon/pkg/render"
	"net/http"
)

var reviews = []models.Review{
	{Name: "Felix Rafael", Rating: "5", Comment: "Salonnya bagus banget aku bisa jadi dirt block dengan pd 😘"},
	{Name: "Ady Wijaya", Rating: "0", Comment: "Jelekk banget aku jadi KEK TELOR mukanya -100/10 😡"},
	{Name: "Hans Christian", Rating: "4", Comment: "Bagus banget aku di rekomendasiin sama pelik makash beb 😍"},
}

func Home(w http.ResponseWriter, r *http.Request) {
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
		// Save the review (you can add code here to save it to a database or file)
		fmt.Printf("New Review: %+v\n", review)

		reviews = append(reviews, review)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
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
	return ""
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
