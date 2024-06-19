package models

// Review is a model for the user review and comment
type Review struct {
	Name    string `json:"name"`
	Rating  string `json:"rating"`
	Comment string `json:"comment"`
}
